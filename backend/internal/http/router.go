package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/YASSERRMD/Amanora/backend/internal/catalog"
	"github.com/YASSERRMD/Amanora/backend/internal/classifier"
	"github.com/YASSERRMD/Amanora/backend/internal/config"
	"github.com/YASSERRMD/Amanora/backend/internal/datasource"
	"github.com/YASSERRMD/Amanora/backend/internal/discovery"
	"github.com/YASSERRMD/Amanora/backend/internal/graph"
	"github.com/YASSERRMD/Amanora/backend/internal/lineage"
	"github.com/YASSERRMD/Amanora/backend/internal/policy"
)

type healthResponse struct {
	Service     string `json:"service"`
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Timestamp   string `json:"timestamp"`
}

func NewRouter(cfg config.Config) http.Handler {
	registry, err := datasource.DefaultRegistry()
	if err != nil {
		panic(err)
	}
	dataSources := datasource.NewService(registry)
	return NewRouterWithServices(
		cfg,
		dataSources,
		discovery.NewService(dataSources, &discovery.MemoryEventPublisher{}, &discovery.MemoryAuditSink{}),
		classifier.NewDefaultService(),
		catalog.NewService(),
		lineage.NewService(graph.NewMemoryRepository()),
		policy.NewService(),
	)
}

func NewRouterWithServices(cfg config.Config, dataSources *datasource.Service, discoveryJobs *discovery.Service, classificationJobs *classifier.Service, catalogService *catalog.Service, lineageService *lineage.Service, policyService *policy.Service) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthHandler(cfg))
	mux.HandleFunc("POST /api/v1/datasources", createDataSourceHandler(dataSources))
	mux.HandleFunc("GET /api/v1/datasources", listDataSourcesHandler(dataSources))
	mux.HandleFunc("GET /api/v1/datasources/{id}", getDataSourceHandler(dataSources))
	mux.HandleFunc("POST /api/v1/datasources/{id}/test", testExistingDataSourceHandler(dataSources))
	mux.HandleFunc("POST /api/v1/datasources/test", testDataSourceHandler(dataSources))
	mux.HandleFunc("POST /api/v1/discovery/jobs", createDiscoveryJobHandler(discoveryJobs))
	mux.HandleFunc("GET /api/v1/discovery/jobs", listDiscoveryJobsHandler(discoveryJobs))
	mux.HandleFunc("GET /api/v1/discovery/jobs/{id}", getDiscoveryJobHandler(discoveryJobs))
	mux.HandleFunc("POST /api/v1/discovery/jobs/{id}/run", runDiscoveryJobHandler(discoveryJobs))
	mux.HandleFunc("POST /api/v1/classification/jobs", createClassificationJobHandler(classificationJobs))
	mux.HandleFunc("GET /api/v1/classification/jobs", listClassificationJobsHandler(classificationJobs))
	mux.HandleFunc("GET /api/v1/classification/findings", listClassificationFindingsHandler(classificationJobs))
	mux.HandleFunc("POST /api/v1/classification/run", runClassificationHandler(classificationJobs))
	mux.HandleFunc("GET /api/v1/catalog/assets", listCatalogAssetsHandler(catalogService))
	mux.HandleFunc("GET /api/v1/catalog/search", listCatalogAssetsHandler(catalogService))
	mux.HandleFunc("GET /api/v1/catalog/assets/{id}", getCatalogAssetHandler(catalogService))
	mux.HandleFunc("GET /api/v1/catalog/fields/{id}", getCatalogFieldHandler(catalogService))
	mux.HandleFunc("POST /api/v1/catalog/assets/{id}/tags", assignCatalogTagsHandler(catalogService))
	mux.HandleFunc("GET /api/v1/lineage/assets/{id}", getLineageHandler(lineageService))
	mux.HandleFunc("GET /api/v1/lineage/assets/{id}/upstream", getLineageUpstreamHandler(lineageService))
	mux.HandleFunc("GET /api/v1/lineage/assets/{id}/downstream", getLineageDownstreamHandler(lineageService))
	mux.HandleFunc("POST /api/v1/lineage/edges", createLineageEdgeHandler(lineageService))
	mux.HandleFunc("POST /api/v1/policies", createPolicyHandler(policyService))
	mux.HandleFunc("GET /api/v1/policies", listPoliciesHandler(policyService))
	mux.HandleFunc("POST /api/v1/policies/{id}/evaluate", evaluatePolicyHandler(policyService))
	mux.HandleFunc("POST /api/v1/policies/evaluate-all", evaluateAllPoliciesHandler(policyService))
	mux.HandleFunc("/", notFoundHandler)

	return mux
}

func createPolicyHandler(policyService *policy.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var doc policy.Document
		if err := decodeJSON(r, &doc); err != nil {
			writeError(w, http.StatusBadRequest, "invalid_json", err)
			return
		}
		created, err := policyService.AddPolicy(r.Context(), doc)
		if err != nil {
			writeError(w, http.StatusBadRequest, "policy_invalid", err)
			return
		}
		writeJSON(w, http.StatusCreated, created)
	}
}

func listPoliciesHandler(policyService *policy.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"items": policyService.ListPolicies()})
	}
}

func evaluatePolicyHandler(policyService *policy.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var facts map[string]string
		if err := decodeJSON(r, &facts); err != nil {
			writeError(w, http.StatusBadRequest, "invalid_json", err)
			return
		}
		decision, err := policyService.Evaluate(r.Context(), r.PathValue("id"), facts)
		if err != nil {
			writeError(w, http.StatusNotFound, "policy_not_found", err)
			return
		}
		writeJSON(w, http.StatusOK, decision)
	}
}

func evaluateAllPoliciesHandler(policyService *policy.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var facts map[string]string
		if err := decodeJSON(r, &facts); err != nil {
			writeError(w, http.StatusBadRequest, "invalid_json", err)
			return
		}
		writeJSON(w, http.StatusOK, policyService.Execute(r.Context(), facts))
	}
}

func getLineageHandler(lineageService *lineage.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := lineageService.AssetLineage(r.Context(), r.PathValue("id"))
		if err != nil {
			writeError(w, http.StatusBadRequest, "lineage_error", err)
			return
		}
		writeJSON(w, http.StatusOK, result)
	}
}

func getLineageUpstreamHandler(lineageService *lineage.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		edges, err := lineageService.Upstream(r.Context(), r.PathValue("id"))
		if err != nil {
			writeError(w, http.StatusBadRequest, "lineage_error", err)
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"items": edges})
	}
}

func getLineageDownstreamHandler(lineageService *lineage.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		edges, err := lineageService.Downstream(r.Context(), r.PathValue("id"))
		if err != nil {
			writeError(w, http.StatusBadRequest, "lineage_error", err)
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"items": edges})
	}
}

func createLineageEdgeHandler(lineageService *lineage.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var edge graph.Edge
		if err := decodeJSON(r, &edge); err != nil {
			writeError(w, http.StatusBadRequest, "invalid_json", err)
			return
		}
		created, err := lineageService.UpsertEdge(r.Context(), edge)
		if err != nil {
			writeError(w, http.StatusBadRequest, "lineage_edge_invalid", err)
			return
		}
		writeJSON(w, http.StatusCreated, created)
	}
}

func listCatalogAssetsHandler(catalogService *catalog.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filter := catalog.SearchFilter{
			Query:          r.URL.Query().Get("q"),
			Classification: r.URL.Query().Get("classification"),
			Owner:          r.URL.Query().Get("owner"),
			Risk:           r.URL.Query().Get("risk"),
		}
		writeJSON(w, http.StatusOK, map[string]any{"items": catalogService.ListAssets(filter)})
	}
}

func getCatalogAssetHandler(catalogService *catalog.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		asset, ok := catalogService.GetAssetDetail(r.PathValue("id"))
		if !ok {
			writeError(w, http.StatusNotFound, "asset_not_found", nil)
			return
		}
		writeJSON(w, http.StatusOK, asset)
	}
}

func getCatalogFieldHandler(catalogService *catalog.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		field, ok := catalogService.GetField(r.PathValue("id"))
		if !ok {
			writeError(w, http.StatusNotFound, "field_not_found", nil)
			return
		}
		writeJSON(w, http.StatusOK, field)
	}
}

func assignCatalogTagsHandler(catalogService *catalog.Service) http.HandlerFunc {
	type request struct {
		Tags []string `json:"tags"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var body request
		if err := decodeJSON(r, &body); err != nil {
			writeError(w, http.StatusBadRequest, "invalid_json", err)
			return
		}
		asset, err := catalogService.AssignTags(r.PathValue("id"), body.Tags)
		if err != nil {
			writeError(w, http.StatusNotFound, "asset_not_found", err)
			return
		}
		writeJSON(w, http.StatusOK, asset)
	}
}

func createClassificationJobHandler(classificationJobs *classifier.Service) http.HandlerFunc {
	type request struct {
		Assets []datasource.AssetMetadata `json:"assets"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var body request
		if err := decodeJSON(r, &body); err != nil {
			writeError(w, http.StatusBadRequest, "invalid_json", err)
			return
		}
		writeJSON(w, http.StatusCreated, classificationJobs.CreateJob(body.Assets))
	}
}

func listClassificationJobsHandler(classificationJobs *classifier.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"items": classificationJobs.ListJobs()})
	}
}

func listClassificationFindingsHandler(classificationJobs *classifier.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"items": classificationJobs.ListFindings()})
	}
}

func runClassificationHandler(classificationJobs *classifier.Service) http.HandlerFunc {
	type request struct {
		JobID  string                     `json:"jobId"`
		Assets []datasource.AssetMetadata `json:"assets"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var body request
		if err := decodeJSON(r, &body); err != nil {
			writeError(w, http.StatusBadRequest, "invalid_json", err)
			return
		}
		jobID := body.JobID
		if jobID == "" {
			jobID = classificationJobs.CreateJob(body.Assets).ID
		}
		job, err := classificationJobs.RunJob(r.Context(), jobID)
		if err != nil {
			writeError(w, http.StatusBadRequest, "classification_job_invalid", err)
			return
		}
		writeJSON(w, http.StatusOK, job)
	}
}

func createDiscoveryJobHandler(discoveryJobs *discovery.Service) http.HandlerFunc {
	type request struct {
		DataSourceID string `json:"dataSourceId"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var body request
		if err := decodeJSON(r, &body); err != nil {
			writeError(w, http.StatusBadRequest, "invalid_json", err)
			return
		}

		job, err := discoveryJobs.CreateJob(r.Context(), body.DataSourceID)
		if err != nil {
			writeError(w, http.StatusBadRequest, "discovery_job_invalid", err)
			return
		}

		writeJSON(w, http.StatusCreated, job)
	}
}

func listDiscoveryJobsHandler(discoveryJobs *discovery.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"items": discoveryJobs.ListJobs()})
	}
}

func getDiscoveryJobHandler(discoveryJobs *discovery.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		job, ok := discoveryJobs.GetJob(r.PathValue("id"))
		if !ok {
			writeError(w, http.StatusNotFound, "discovery_job_not_found", nil)
			return
		}

		writeJSON(w, http.StatusOK, job)
	}
}

func runDiscoveryJobHandler(discoveryJobs *discovery.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		job, err := discoveryJobs.RunJob(r.Context(), r.PathValue("id"))
		if err != nil {
			writeJSON(w, http.StatusBadRequest, job)
			return
		}

		writeJSON(w, http.StatusOK, job)
	}
}

func createDataSourceHandler(dataSources *datasource.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var source datasource.DataSource
		if err := decodeJSON(r, &source); err != nil {
			writeError(w, http.StatusBadRequest, "invalid_json", err)
			return
		}

		created, err := dataSources.Register(source)
		if err != nil {
			writeError(w, http.StatusBadRequest, "data_source_invalid", err)
			return
		}

		writeJSON(w, http.StatusCreated, created)
	}
}

func listDataSourcesHandler(dataSources *datasource.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"items": dataSources.List()})
	}
}

func getDataSourceHandler(dataSources *datasource.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		source, ok := dataSources.Get(r.PathValue("id"))
		if !ok {
			writeError(w, http.StatusNotFound, "data_source_not_found", nil)
			return
		}

		writeJSON(w, http.StatusOK, source)
	}
}

func testExistingDataSourceHandler(dataSources *datasource.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := dataSources.TestDataSource(r.Context(), r.PathValue("id"))
		if err != nil {
			writeJSON(w, http.StatusNotFound, result)
			return
		}

		writeJSON(w, http.StatusOK, result)
	}
}

func testDataSourceHandler(dataSources *datasource.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cfg datasource.ConnectionConfig
		if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil {
			writeError(w, http.StatusBadRequest, "invalid_json", err)
			return
		}

		result, err := dataSources.TestConnection(r.Context(), cfg)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, result)
			return
		}

		writeJSON(w, http.StatusOK, result)
	}
}

func healthHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, healthResponse{
			Service:     "amanora-api",
			Status:      "ok",
			Environment: cfg.Environment,
			Timestamp:   time.Now().UTC().Format(time.RFC3339),
		})
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotFound, map[string]string{
		"error":   "not_found",
		"service": "amanora-api",
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, code string, err error) {
	message := code
	if err != nil {
		message = err.Error()
	}
	message = strings.TrimSpace(message)
	if message == "" {
		message = code
	}
	writeJSON(w, status, map[string]string{
		"error":   code,
		"message": message,
	})
}

func decodeJSON(r *http.Request, target any) error {
	if r.Body == nil {
		return errors.New("request body is required")
	}
	return json.NewDecoder(r.Body).Decode(target)
}

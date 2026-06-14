package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/YASSERRMD/Amanora/backend/internal/config"
	"github.com/YASSERRMD/Amanora/backend/internal/datasource"
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
	return NewRouterWithServices(cfg, datasource.NewService(registry))
}

func NewRouterWithServices(cfg config.Config, dataSources *datasource.Service) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthHandler(cfg))
	mux.HandleFunc("POST /api/v1/datasources", createDataSourceHandler(dataSources))
	mux.HandleFunc("GET /api/v1/datasources", listDataSourcesHandler(dataSources))
	mux.HandleFunc("GET /api/v1/datasources/{id}", getDataSourceHandler(dataSources))
	mux.HandleFunc("POST /api/v1/datasources/{id}/test", testExistingDataSourceHandler(dataSources))
	mux.HandleFunc("POST /api/v1/datasources/test", testDataSourceHandler(dataSources))
	mux.HandleFunc("/", notFoundHandler)

	return mux
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

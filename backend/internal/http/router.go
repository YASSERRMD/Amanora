package httpapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/YASSERRMD/Amanora/backend/internal/config"
)

type healthResponse struct {
	Service     string `json:"service"`
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Timestamp   string `json:"timestamp"`
}

func NewRouter(cfg config.Config) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthHandler(cfg))
	mux.HandleFunc("/", notFoundHandler)

	return mux
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

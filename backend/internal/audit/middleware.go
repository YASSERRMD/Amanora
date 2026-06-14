package audit

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func Middleware(next http.Handler) http.Handler {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		next.ServeHTTP(w, r)
		logger.Info("request audited", "method", r.Method, "path", r.URL.Path, "duration_ms", time.Since(started).Milliseconds())
	})
}

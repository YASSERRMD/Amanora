package httpapi

import (
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/YASSERRMD/Amanora/backend/internal/audit"
	"github.com/YASSERRMD/Amanora/backend/internal/tenant"
)

func secure(handler http.Handler) http.Handler {
	return audit.Middleware(tenant.Middleware(rateLimit(apiKeyAuth(secureHeaders(handler)))))
}

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "no-referrer")
		next.ServeHTTP(w, r)
	})
}

func apiKeyAuth(next http.Handler) http.Handler {
	required := os.Getenv("AMANORA_API_KEY")
	if required == "" {
		required = "dev-api-key"
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/healthz" {
			next.ServeHTTP(w, r)
			return
		}
		if r.Header.Get("X-API-Key") != required {
			writeError(w, http.StatusUnauthorized, "unauthorized", nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func rateLimit(next http.Handler) http.Handler {
	var mu sync.Mutex
	counts := map[string]int{}
	windowStarted := time.Now()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		if time.Since(windowStarted) > time.Minute {
			counts = map[string]int{}
			windowStarted = time.Now()
		}
		key := r.RemoteAddr
		counts[key]++
		allowed := counts[key] <= 600
		mu.Unlock()
		if !allowed {
			writeError(w, http.StatusTooManyRequests, "rate_limited", nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func redactSensitive(value string) string {
	lower := strings.ToLower(value)
	if strings.Contains(lower, "password") || strings.Contains(lower, "secret") || strings.Contains(lower, "token") {
		return "[redacted]"
	}
	return value
}

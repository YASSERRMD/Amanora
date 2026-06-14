package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecurityMiddlewareAllowsHealthAndProtectsAPI(t *testing.T) {
	handler := secure(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	health := httptest.NewRecorder()
	handler.ServeHTTP(health, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	if health.Code != http.StatusOK {
		t.Fatalf("expected health to pass, got %d", health.Code)
	}
	if health.Header().Get("X-Content-Type-Options") != "nosniff" {
		t.Fatal("expected secure header")
	}

	blocked := httptest.NewRecorder()
	handler.ServeHTTP(blocked, httptest.NewRequest(http.MethodGet, "/api/v1/catalog/assets", nil))
	if blocked.Code != http.StatusUnauthorized {
		t.Fatalf("expected unauthorized, got %d", blocked.Code)
	}

	allowedReq := httptest.NewRequest(http.MethodGet, "/api/v1/catalog/assets", nil)
	allowedReq.Header.Set("X-API-Key", "dev-api-key")
	allowed := httptest.NewRecorder()
	handler.ServeHTTP(allowed, allowedReq)
	if allowed.Code != http.StatusOK {
		t.Fatalf("expected allowed request, got %d", allowed.Code)
	}
}

func TestRedactSensitive(t *testing.T) {
	if redactSensitive("db_password=secret") != "[redacted]" {
		t.Fatal("expected redaction")
	}
	if redactSensitive("ordinary") != "ordinary" {
		t.Fatal("expected ordinary value unchanged")
	}
}

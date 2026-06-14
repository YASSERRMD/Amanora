package tenant

import (
	"context"
	"net/http"
)

type contextKey string

const tenantKey contextKey = "tenant_id"

func WithTenant(ctx context.Context, tenantID string) context.Context {
	return context.WithValue(ctx, tenantKey, tenantID)
}

func FromContext(ctx context.Context) string {
	value, _ := ctx.Value(tenantKey).(string)
	return value
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenantID := r.Header.Get("X-Tenant-ID")
		if tenantID == "" {
			tenantID = "default"
		}
		next.ServeHTTP(w, r.WithContext(WithTenant(r.Context(), tenantID)))
	})
}

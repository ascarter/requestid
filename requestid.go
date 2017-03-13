package requestid

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxKey int

const ridKey ctxKey = ctxKey(0)

// NewContext creates a context with request id
func NewContext(ctx context.Context, rid string) context.Context {
	return context.WithValue(ctx, ridKey, rid)
}

// FromContext returns the request id from context
func FromContext(ctx context.Context) (string, bool) {
	rid, ok := ctx.Value(ridKey).(string)
	return rid, ok
}

// RequestIDHandler sets unique request id.
// If header `X-Request-ID` is already present in the request, that is considered the
// request id. Otherwise, generates a new unique ID.
func RequestIDHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rid := r.Header.Get("X-Request-ID")
		if rid == "" {
			rid = uuid.New().String()
			r.Header.Set("X-Request-ID", rid)
		}
		ctx := NewContext(r.Context(), rid)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

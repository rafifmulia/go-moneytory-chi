package middleware

import (
	"net/http"
	"restfulapi/helper"
)

// Authorization Middleware
func AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != "secret" {
			helper.RespUnauthorized(w, "")
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

package middleware

import "net/http"

func (mw *MiddlewareManager) CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set allowed origins (replace "*" with specific origins if needed)
		w.Header().Set("Access-Control-Allow-Origin", "*") // Adjust for specific origins

		// Set allowed methods (adjust as required)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE ")

		// Set allowed headers (adjust as required)
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, X-Request-Id, X-CSRF-Token")

		if r.Method != "OPTIONS" {
			next.ServeHTTP(w, r)
			return
		}

		// Handle CORS preflight requests
		w.WriteHeader(http.StatusOK)
	})
}

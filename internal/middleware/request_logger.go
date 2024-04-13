package middleware

import (
	"net/http"
	"time"
)

type responseLogger struct {
	http.ResponseWriter
	statusCode int
}

func (rl *responseLogger) WriteHeader(code int) {
	rl.statusCode = code
	rl.ResponseWriter.WriteHeader(code)
}

func (rl *responseLogger) Write(b []byte) (int, error) {
	// Log the response here if needed
	return rl.ResponseWriter.Write(b)
}

func (mw *MiddlewareManager) RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rl := &responseLogger{ResponseWriter: w, statusCode: http.StatusOK}
		defer func() {
			s := time.Since(start).String()
			mw.logger.Infof(" Method: %s, RequestURI: %v, Size: %v, Time: %s StatusCode: %d",
				r.Method, r.RequestURI, r.ContentLength, s, rl.statusCode)
		}()

		next.ServeHTTP(rl, r)
	})
}

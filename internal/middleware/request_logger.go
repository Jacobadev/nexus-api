package middleware

import (
	"net/http"
	"time"

	"github.com/gateway-address/pkg/utils"
)

func (mw *MiddlewareManager) RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			s := time.Since(start).String()
			ipAddress := utils.GetIPAddress(r)
			mw.logger.Infof(" Method: %s, RequestURI: %v, Size: %v, Time: %s, IPAddress: %s, ",
				r.Method, r.RequestURI, r.ContentLength, s, ipAddress)
		}()

		next.ServeHTTP(w, r)
	})
}

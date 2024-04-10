package middleware

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func (mw *MiddlewareManager) DebugMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if mw.cfg.Server.Debug {
			dump, err := httputil.DumpRequest(r, true)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			mw.logger.Info(fmt.Sprintf("\nRequest dump begin :--------------\n\n%s\n\nRequest dump end :--------------", dump))
		}
		next.ServeHTTP(w, r)
	})
}

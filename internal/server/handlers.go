package server

import (
	"github.com/gateway-address/internal/auth/repository"
	"github.com/gateway-address/internal/auth/usecase"
	"github.com/gateway-address/internal/delivery/http"
	"github.com/gorilla/mux"
)

// Map Server Handlers
func (s *Server) MapHandlers(r *mux.Router) error {
	// Init repositories
	authRepo := repository.NewRepositorySqlite(s.db)

	authUC := usecase.NewAuthUseCase(s.cfg, authRepo, s.logger)
	authHandler := http.NewAuthHandlers(s.cfg, authUC, s.logger)
	// Init handlers
	// mw := .NewMiddlewareManager(authUC, s.cfg, []string{"*"}, s.logger)
	http.HealthCheck(r.PathPrefix("/health").Subrouter())

	v1 := r.PathPrefix("/api/v1/auth").Subrouter()

	http.MapAuthRoutes(authHandler, v1)

	return nil
}

package server

import (
	"github.com/gateway-address/internal/auth/repository"
	"github.com/gateway-address/internal/auth/usecase"
	"github.com/gateway-address/internal/delivery/http"
)

// Map Server Handlers
func (s *Server) MapHandlers() error {
	authRepo := repository.NewRepositorySqlite(s.db)
	authUC := usecase.NewAuthUseCase(s.cfg, authRepo, s.logger)
	authHandler := http.NewAuthHandlers(s.cfg, authUC, s.logger)
	// mw := .NewMiddlewareManager(authUC, s.cfg, []string{"*"}, s.logger)

	v1 := s.router.PathPrefix("/api/v1/auth").Subrouter()
	http.MapAuthRoutes(authHandler, v1)

	return nil
}

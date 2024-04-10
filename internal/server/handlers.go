package server

import (
	"github.com/gateway-address/internal/auth/repository"
	authUseCase "github.com/gateway-address/internal/auth/usecase"
	"github.com/gateway-address/internal/delivery/http"
	"github.com/gateway-address/internal/middleware"
	sessRepository "github.com/gateway-address/internal/session/repository"
	sessUseCase "github.com/gateway-address/internal/session/usecase"
)

// Map Server Handlers
func (s *Server) MapHandlers() error {
	authRepo := repository.NewRepositorySqlite(s.db)
	authUC := authUseCase.NewAuthUseCase(s.cfg, authRepo, s.logger)
	authHandler := http.NewAuthHandlers(s.cfg, authUC, s.logger)

	sessRepo := sessRepository.NewSessionRepository(s.redisClient, s.cfg)
	sessUC := sessUseCase.NewSessionUseCase(sessRepo, s.cfg)
	mw := middleware.NewMiddlewareManager(sessUC, authUC, s.cfg, []string{"*"}, s.logger)
	s.router.Use(mw.RequestLoggerMiddleware)
	s.router.Use(mw.DebugMiddleware)
	apiV1 := s.router.PathPrefix("/api/v1").Subrouter()

	// Define subrouters for different sections of the API
	authGroup := apiV1.PathPrefix("/auth").Subrouter()
	http.MapAuthRoutes(authHandler, authGroup) // Added closing parenthesis
	http.HealthCheck(apiV1)
	return nil
}

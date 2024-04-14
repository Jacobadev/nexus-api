package server

import (
	"github.com/gateway-address/internal/auth/delivery/http"
	"github.com/gateway-address/internal/auth/repository"
	authUseCase "github.com/gateway-address/internal/auth/usecase"
	"github.com/gateway-address/internal/middleware"
	sessRepository "github.com/gateway-address/internal/session/repository"
	sessUseCase "github.com/gateway-address/internal/session/usecase"
	"github.com/gateway-address/internal/websocket"
	websocketHttp "github.com/gateway-address/internal/websocket/delivery/http"
)

// Map Server Handlers
func (s *Server) MapHandlers() error {
	sessRepo := sessRepository.NewSessionRepository(s.redisClient, s.cfg)
	sessUC := sessUseCase.NewSessionUseCase(sessRepo, s.cfg)

	authRepo := repository.NewRepositorySqlite(s.db)
	authUC := authUseCase.NewAuthUseCase(s.cfg, authRepo, s.logger)
	authHandler := http.NewAuthHandlers(s.cfg, sessUC, authUC, s.logger)

	wsUC := websocket.NewUseCase(s.cfg, s.logger)
	wsHandler := websocketHttp.NewWebSocket(s.cfg, wsUC, s.logger)

	mw := middleware.NewMiddlewareManager(sessUC, authUC, s.cfg, []string{"*"}, s.logger)
	s.router.Use(mw.RequestLoggerMiddleware)

	// s.router.Use(mw.DebugMiddleware)
	apiV1 := s.router.PathPrefix("/api/v1").Subrouter()

	// Define subrouters for different sections of the API
	authGroup := apiV1.PathPrefix("/auth").Subrouter()
	wsGroup := apiV1.PathPrefix("/ws").Subrouter()

	http.MapAuthRoutes(authHandler, authGroup) // Added closing parenthesis
	websocketHttp.MapWebSocketRoutes(wsGroup)
	http.HealthCheck(apiV1)
	return nil
}

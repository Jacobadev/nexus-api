package middleware

import (
	"github.com/gateway-address/config"
	"github.com/gateway-address/internal/auth"
	"github.com/gateway-address/internal/session"
	"github.com/gateway-address/pkg/logger"
)

// Middleware manager
type MiddlewareManager struct {
	sessUC  session.UCSession
	authUC  auth.UseCase
	cfg     *config.Config
	origins []string
	logger  logger.Logger
}

// Middleware manager constructor
func NewMiddlewareManager(sessUC session.UCSession, authUC auth.UseCase, cfg *config.Config, origins []string, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{sessUC: sessUC, authUC: authUC, cfg: cfg, origins: origins, logger: logger}
}

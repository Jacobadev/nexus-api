package http

import (
	"net/http"

	"github.com/gateway-address/config"
	"github.com/gateway-address/internal/websocket"
	"github.com/gateway-address/pkg/logger"
)

type websocketHandlers struct {
	cfg         *config.Config
	websocketUC websocket.UseCase
	logger      logger.Logger
}

func NewWebSocket(cfg *config.Config, websocketUC websocket.UseCase, logger logger.Logger) *websocketHandlers {
	return &websocketHandlers{
		cfg:         cfg,
		websocketUC: websocketUC,
		logger:      logger,
	}
}

func (w *websocketHandlers) changeLPEvent() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	}
}

func (w *websocketHandlers) newOrderEvent() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	}
}

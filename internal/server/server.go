package server

import (
	"net/http"
	"time"

	"github.com/gateway-address/config"
	"github.com/gateway-address/pkg/logger"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	router      *mux.Router
	cfg         *config.Config
	db          *sqlx.DB
	logger      logger.Logger
	redisClient *redis.Client
}

const (
	ctxTimeout     = 5
	maxHeaderBytes = 1 << 20
)

func NewServer(cfg *config.Config, db *sqlx.DB, logger logger.Logger) *Server {
	return &Server{cfg: cfg, router: mux.NewRouter(), db: db, logger: logger}
}

func (s *Server) Run() error {
	if err := s.MapHandlers(s.router); err != nil {
		return err
	}
	server := &http.Server{
		Addr:         "0.0.0.0:" + s.cfg.Server.Port,
		Handler:      s.router,
		ReadTimeout:  time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout: time.Second * s.cfg.Server.WriteTimeout,
	}
	s.logger.Infof("Server is listening on: %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		s.logger.Infof("Error starting server: %s", err)
	}

	return nil
}

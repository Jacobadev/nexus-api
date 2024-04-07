package main

import (
	"log"
	"net/http"

	"github.com/gateway-address/config"
	"github.com/gateway-address/internal/server"
	"github.com/gateway-address/pkg/db/postgres"
	"github.com/gateway-address/pkg/logger"
	"github.com/gateway-address/pkg/utils"
	"github.com/gorilla/mux"
)

func HealthCheck(router *mux.Router) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}

func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath()

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)

	// redisClient := redis.NewRedisClient(cfg)
	// defer redisClient.Close()
	db, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	s := server.NewServer(cfg, db, appLogger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}

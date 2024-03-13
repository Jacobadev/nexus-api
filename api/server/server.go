package server

import (
	"fmt"
	"net/http"

	"github.com/gateway-address/api/config"
	"github.com/gateway-address/api/routes"
)

func serverGetPort() (int, error) {
	cfgFile, err := config.LoadConfig()
	if err != nil {
		return 0, err
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		return 0, err
	}
	return cfg.Server.Port, nil
}

func Start(mux *http.ServeMux) {
	port, err := serverGetPort()
	if err != nil {
		fmt.Printf("%s", err)
	}
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("Listening on %s\n", address)

	err = http.ListenAndServe(address, mux)
	if err != nil {
		panic(err)
	}
}

func RegisterUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/user", routes.UserMethodController)
}

func GetMuxV1() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}

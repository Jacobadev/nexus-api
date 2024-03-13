package main

import (
	"fmt"
	"net/http"

	"github.com/gateway-address/api/config"
	"github.com/gateway-address/api/routes"
)

func RegisterUserRoutes() {
	http.HandleFunc("/user", routes.UserMethodController)
}

func main() {
	port, err := serverGetPort()
	if err != nil {
		fmt.Printf("%s", err)
	}
	http.HandleFunc("/", "<h1>Hello World</h1>")
	start(port)
}

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

func start(port int) {
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("Listening on %s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/gateway-address/routes"
	"github.com/gateway-address/server"
)

func main() {
	mux := server.GetMuxRouterV1()
	routes.RegisterUserHandler(mux)
	server.StartServer(mux)
}

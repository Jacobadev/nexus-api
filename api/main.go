package main

import (
	"github.com/gateway-address/server"
)

func main() {
	mux := server.GetMuxV1()
	server.RegisterRoutes(mux)
	server.StartServer(mux)
}

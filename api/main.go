package main

import (
	"github.com/gateway-address/api/server"
)

func main() {
	mux := server.GetMuxV1()
	server.RegisterUserRoutes(mux)
	server.Start(mux)
}

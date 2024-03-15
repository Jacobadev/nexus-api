package main

import (
	"github.com/gateway-address/handler"
	"github.com/gateway-address/server"
	"github.com/gorilla/mux"
)

func RegisterUserHandler(mux *mux.Router) {
	mux.HandleFunc("/user", handler.HandleGetAllUsers).Methods("GET")
	mux.HandleFunc("/user", handler.HandleCreateUser).Methods("POST")
}

func main() {
	mux := server.GetMuxRouterV1()
	RegisterUserHandler(mux)
	server.StartServer(mux)
}

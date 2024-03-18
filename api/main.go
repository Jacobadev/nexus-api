package main

import (
	"fmt"

	"github.com/gateway-address/handler"
	"github.com/gateway-address/repository"
	"github.com/gateway-address/server"
	"github.com/gateway-address/user"
	"github.com/gorilla/mux"
)

func RegisterUserHandler(router *mux.Router, userHandler *handler.UserHandler) {
	router.HandleFunc("/user", handler.CreateUserHandler(userHandler)).Methods("POST")
	router.HandleFunc("/user", handler.GetUsersHandler(userHandler)).Methods("GET")
	router.HandleFunc("/user/limit={limit}/offset={offset}", handler.GetPaginatedUsersHandler(userHandler)).Methods("GET")

	router.HandleFunc("/user/{id}", handler.GetUserByIDHandler(userHandler)).Methods("GET")

	// router.HandleFunc("/user/{limit}/{offset}", handler.GetPaginatedUsersHandler(userHandler)).Methods("GET")
}

func main() {
	mux := server.GetMuxSubRouterV1()

	var userRepository user.UserRepository
	repo, err := repository.NewRepositorySqlite()

	userHandler := handler.NewUserHandler(userRepository, repo)
	if err != nil {
		fmt.Println(err)
	}

	RegisterUserHandler(mux, userHandler)
	server.StartServer(mux)
}

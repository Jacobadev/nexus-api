package main

import (
	"fmt"
	"net/http"

	"github.com/gateway-address/handler"
	"github.com/gateway-address/repository"
	"github.com/gateway-address/server"
	"github.com/gateway-address/user"
	"github.com/gorilla/mux"
)

func RegisterUserHandler(router *mux.Router, userHandler *handler.UserHandler) {
	router.HandleFunc("/user", handler.CreateUserHandler(userHandler)).Methods("POST")
	router.HandleFunc("/user", handler.GetUsersHandler(userHandler)).Methods("GET")
	router.HandleFunc("/user/{id}", handler.GetUserByIDHandler(userHandler)).Methods("GET")
	router.HandleFunc("/user/{id}", handler.DeleteUserByIDHandler(userHandler)).Methods("DELETE")
	router.HandleFunc("/user/{id}", handler.UpdateUserByIDHandler(userHandler)).Methods("PUT")
	router.HandleFunc("/user/{id}", handler.PartialUpdateUserByIDHandler(userHandler)).Methods("PATCH")
}

func RootHandler(router *mux.Router) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status": "OK"}`)
	}).Methods("GET")
}

func main() {
	mux := server.GetMuxSubRouterV1()

	var userRepository user.UserRepository
	repo, err := repository.NewRepositorySqlite()

	userHandler := handler.NewUserHandler(userRepository, repo)
	if err != nil {
		fmt.Println(err)
	}

	RootHandler(mux)
	RegisterUserHandler(mux, userHandler)
	server.StartServer(mux)
}

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
	router.HandleFunc("/", handler.CreateUserHandler(userHandler)).Methods("POST")
	router.HandleFunc("/", handler.GetUsersHandler(userHandler)).Methods("GET")
	router.HandleFunc("{id}", handler.GetUserByIDHandler(userHandler)).Methods("GET")
	router.HandleFunc("/{id}", handler.DeleteUserByIDHandler(userHandler)).Methods("DELETE")
	router.HandleFunc("/{id}", handler.UpdateUserByIDHandler(userHandler)).Methods("PUT")
	router.HandleFunc("/{id}", handler.PartialUpdateUserByIDHandler(userHandler)).Methods("PATCH")
}

func RootHandler(router *mux.Router) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status": "OK"}`)
	}).Methods("GET")
}

func main() {
	mux := server.GetUserRouter()

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

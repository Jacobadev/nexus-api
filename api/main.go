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

func RegisterUserHandler(mux *mux.Router, userHandler *handler.UserHandler) {
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		handler.GetUsersHandler(userHandler)(w, r)
	}).Methods("GET")

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateUserHandler(userHandler)(w, r)
		userData := handler.ExtractUserInput(r)

		// Imprimir os dados do usuário para verificar se estão corretos
		fmt.Println("Dados do usuário recebidos:")
		fmt.Printf("First Name: %s\n", userData.FirstName)
		fmt.Printf("Last Name: %s\n", userData.LastName)
		fmt.Printf("Username: %s\n", userData.UserName)
		fmt.Printf("Email: %s\n", userData.Email)
		fmt.Printf("Password: %s\n", userData.Password)
	}).Methods("POST")
}

func RegisterUserPOSTHandler(mux *mux.Router) {
}

func main() {
	mux := server.GetMuxRouterV1()

	var userRepository user.UserRepository
	repo, err := repository.NewRepositorySqlite()

	userHandler := handler.NewUserHandler(userRepository, repo)
	if err != nil {
		fmt.Println(err)
	}

	RegisterUserHandler(mux, userHandler)
	server.StartServer(mux)
}

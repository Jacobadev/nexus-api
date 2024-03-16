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
		if r.Method == http.MethodPost {
			// Extrair os dados do usuário do request
			handler.CreateUserHandler(userHandler)(w, r)
			return
		}
		handler.GetUsersHandler(userHandler)(w, r)
	}).Methods("GET", "POST") // Permitir tanto GET quanto POST para a rota "/user"
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

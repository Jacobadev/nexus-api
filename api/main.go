package main

import (
	"net/http"

	"github.com/gateway-address/handler"
	"github.com/gateway-address/repository"
	"github.com/gateway-address/server"
	"github.com/gateway-address/user"
	"github.com/gorilla/mux"
)

func RegisterUserHandler(mux *mux.Router) {
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		var userRepository user.UserRepository
		repositorySqlite, err := repository.NewRepositorySqlite()
		if err != nil {
			// Handle error
			http.Error(w, "Failed to initialize repository", http.StatusInternalServerError)
			return
		}

		handleUser := handler.NewHandleUser(userRepository, repositorySqlite)

		// Call HandleCreateUser function passing handleUser instance
		handler.HandleGetUsers(handleUser)(w, r)
	}).Methods("GET")
}

func main() {
	mux := server.GetMuxRouterV1()
	RegisterUserHandler(mux)
	server.StartServer(mux)
}

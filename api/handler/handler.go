package handler

import (
	"fmt"
	"net/http"

	"github.com/gateway-address/repository"
	"github.com/gateway-address/server"
	"github.com/gateway-address/user"
	"github.com/gateway-address/validate"
)

type HandleUser struct {
	UserRepository user.UserRepository
	Repository     *repository.RepositorySqlite
}

func NewHandleUser(userRepository user.UserRepository, repo *repository.RepositorySqlite) *HandleUser {
	return &HandleUser{
		UserRepository: userRepository,
		Repository:     repo,
	}
}

func HandleCreateUser(handleUser *HandleUser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := ExtractUserInput(r)
		v := validate.NewValidator()
		err := v.ValidateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = handleUser.Repository.Create(user) // Usando o método CreateUser da interface UserRepository
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		server.WriteUserPOSTResponse(w)
	}
}

func HandleGetUsers(handleUser *HandleUser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := handleUser.Repository.GetAll() // Usando o método GetAll da interface UserRepository
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get users: %v", err), http.StatusInternalServerError)
			return
		}

		server.WriteUserGETResponse(w, users)
	}
}

package handler

import (
	"fmt"
	"net/http"

	"github.com/gateway-address/repository"
	"github.com/gateway-address/server"
	"github.com/gateway-address/user"
)

type UserHandler struct {
	UserRepository user.UserRepository
	Repository     *repository.RepositorySqlite
}

func NewUserHandler(userRepository user.UserRepository, repo *repository.RepositorySqlite) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
		Repository:     repo,
	}
}

func CreateUserHandler(userHandler *UserHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := user.ExtractUserInput(r)
		fmt.Println(user.FirstName)
		fmt.Println(user.LastName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// v := validate.NewValidator()
		// err := v.ValidateUser(user)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		//
		err = userHandler.Repository.Create(user)
		if err != nil {
			fmt.Println(err)
			http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
			return
		}

		server.WriteUserPOSTResponse(w)
	}
}

func GetUsersHandler(userHandler *UserHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userHandler.Repository.GetAll()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get users: %v", err), http.StatusInternalServerError)
			return
		}

		server.WriteUserGETResponse(w, users)
	}
}

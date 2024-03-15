package handler

import (
	"fmt"
	"net/http"

	"github.com/gateway-address/repository"
	"github.com/gateway-address/server"
	"github.com/gateway-address/validation"
)

type UserHandler struct {
	user repository.UserRepository
}

func NewUserHandler(user repository.UserRepository) *UserHandler {
	return &UserHandler{user: user}
}

// HandleGetAllUsers lida com a requisição para obter todos os usuários.
func (uc *UserHandler) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.user.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	server.WriteUserGETResponse(w, users)
}

func (uc *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := validation.ValidateUserInput(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = uc.user.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	server.WriteUserPOSTResponse(w)
}

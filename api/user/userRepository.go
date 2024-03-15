package repository

import (
	"github.com/gateway-address/model"
)

type UserOperations interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(userInfo model.User) error
}

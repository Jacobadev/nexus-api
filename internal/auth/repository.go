package auth

import (
	"database/sql"

	model "github.com/gateway-address/internal/models"
)

type Repository interface {
	Register(user model.User) (sql.Result, error)
	GetAll() ([]interface{}, error)
	GetById(id int) (interface{}, error)
	UpdateByID(id int, user model.User) (sql.Result, error)
	PartialUpdateByID(id int, user model.User) (sql.Result, error)
	DeleteByID(id int) (sql.Result, error)
	FindByEmail(*model.User) (*model.User, error)
}

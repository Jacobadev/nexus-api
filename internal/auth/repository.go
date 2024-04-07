package auth

import (
	"database/sql"

	model "github.com/gateway-address/internal/models"
)

type Repository interface {
	GetAll() ([]interface{}, error)
	GetById(id int) (interface{}, error)
	Create(userInfo model.User) (sql.Result, error)
	UpdateByID(id int, userInfo model.User) (sql.Result, error)
	DeleteByID(id int) (sql.Result, error)
	PartialUpdateByID(id int, userInfo model.User) (sql.Result, error)
}

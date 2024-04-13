package auth

import (
	"database/sql"

	model "github.com/gateway-address/internal/models"
	"github.com/gateway-address/pkg/utils"
)

type Repository interface {
	Register(user model.User) (sql.Result, error)
	GetUsers(pq *utils.PaginationQuery) (*model.UsersList, error)
	GetById(id int) (interface{}, error)
	UpdateByID(id int, user model.User) (sql.Result, error)
	PartialUpdateByID(id int, user model.User) (sql.Result, error)
	Delete(id int) (sql.Result, error)
	FindByEmail(*model.User) (*model.User, error)
}

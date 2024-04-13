//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package auth

import (
	model "github.com/gateway-address/internal/models"
	"github.com/gateway-address/pkg/utils"
)

// Auth repository interface
type UseCase interface {
	Register(user *model.User) (*model.UserWithToken, error)
	Login(user *model.User) (*model.UserWithToken, error)
	GetByID(userID int) (*model.User, error)
	Delete(userID int) error
	GetUsers(pq *utils.PaginationQuery) (*model.UsersList, error)
	// Update(ctx context.Context, user *model.User) (*model.User, error)
	// FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*model.UsersList, error)
}

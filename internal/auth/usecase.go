//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package auth

import (
	model "github.com/gateway-address/internal/models"
)

// Auth repository interface
type UseCase interface {
	Register(user *model.User) (*model.UserWithToken, error)
	Login(user *model.User) (*model.UserWithToken, error)
	// Update(ctx context.Context, user *model.User) (*model.User, error)
	// Delete(ctx context.Context, userID uuid.UUID) error
	// GetByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	// FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*model.UsersList, error)
	// GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*model.UsersList, error)
}

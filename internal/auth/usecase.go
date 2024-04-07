//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package auth

// Auth repository interface
type UseCase interface {
	Register()
	// Login(ctx context.Context, user *model.User) error
	// Update(ctx context.Context, user *model.User) (*model.User, error)
	// Delete(ctx context.Context, userID uuid.UUID) error
	// GetByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	// FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*model.UsersList, error)
	// GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*model.UsersList, error)
}

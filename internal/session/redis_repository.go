//go:generate mockgen -source redis_repository.go -destination mock/redis_repository_mock.go -package mock
package session

import (
	"context"

	model "github.com/gateway-address/internal/models"
)

// Session repository
type SessRepository interface {
	CreateSession(ctx context.Context, session *model.Session, expire int) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*model.Session, error)
	DeleteByID(ctx context.Context, sessionID string) error
}

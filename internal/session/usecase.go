//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package session

import (
	"context"

	model "github.com/gateway-address/internal/models"
)

// Session use case
type UCSession interface {
	CreateSession(ctx context.Context, session *model.Session, expire int) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*model.Session, error)
	DeleteByID(ctx context.Context, sessionID string) error
}

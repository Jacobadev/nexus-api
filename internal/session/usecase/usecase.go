package usecase

import (
	"context"

	"github.com/gateway-address/config"
	model "github.com/gateway-address/internal/models"
	"github.com/gateway-address/internal/session"
)

type sessionUC struct {
	sessionRepo session.SessRepository
	cfg         *config.Config
}

// New session use case constructor
func NewSessionUseCase(sessionRepo session.SessRepository, cfg *config.Config) session.UCSession {
	return &sessionUC{sessionRepo: sessionRepo, cfg: cfg}
}

// Create new session
func (u *sessionUC) CreateSession(ctx context.Context, session *model.Session, expire int) (string, error) {
	return u.sessionRepo.CreateSession(ctx, session, expire)
}

// Delete session by id
func (u *sessionUC) DeleteByID(ctx context.Context, sessionID string) error {
	return u.sessionRepo.DeleteByID(ctx, sessionID)
}

// get session by id
func (u *sessionUC) GetSessionByID(ctx context.Context, sessionID string) (*model.Session, error) {
	return u.sessionRepo.GetSessionByID(ctx, sessionID)
}

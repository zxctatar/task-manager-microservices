package login

import (
	"context"
	"log/slog"
	"userservice/internal/repository/hasher"
	"userservice/internal/repository/session"
	storagerepo "userservice/internal/repository/storage"
	logmodel "userservice/internal/usecase/models/login"
)

type LoginUC struct {
	log *slog.Logger

	storage     storagerepo.StorageRepo
	passHasher  hasher.PasswordHasher
	sessionRepo session.SessionRepo
}

func NewLoginUC(log *slog.Logger, storage storagerepo.StorageRepo, passHasher hasher.PasswordHasher, sessionRepo session.SessionRepo) *LoginUC {
	return &LoginUC{
		log:         log,
		storage:     storage,
		passHasher:  passHasher,
		sessionRepo: sessionRepo,
	}
}

func (l *LoginUC) Login(ctx context.Context, in *logmodel.LoginInput) (*logmodel.LoginOutput, error) {
	panic("not implemented")
}

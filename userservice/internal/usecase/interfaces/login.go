package interfaces

import (
	"context"
	logmodel "userservice/internal/usecase/models/login"
)

type LoginUsecase interface {
	Login(ctx context.Context, in *logmodel.LoginInput) (*logmodel.LoginOutput, error)
}

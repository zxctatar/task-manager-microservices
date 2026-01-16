package storagerepo

import (
	"context"
	userdomain "userservice/internal/domain/user"
)

type StorageRepo interface {
	Save(ctx context.Context, ud *userdomain.UserDomain) (uint32, error)
	FindByEmail(ctx context.Context, email string) (*userdomain.UserDomain, error)
}

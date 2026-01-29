package storage

import (
	"context"
	projectdomain "projectservice/internal/domain/project"
)

type StorageRepo interface {
	FindByName(ctx context.Context, ownerId uint32, name string) (*projectdomain.ProjectDomain, error)
	Save(ctx context.Context, proj *projectdomain.ProjectDomain) error
}

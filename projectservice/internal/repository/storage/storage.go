package storage

import (
	"context"
	projectdomain "projectservice/internal/domain/project"
)

type Storage interface {
	Save(ctx context.Context, proj *projectdomain.ProjectDomain) error
	Delete(ctx context.Context, proj *projectdomain.ProjectDomain) error
	GetAll(ctx context.Context, ownerId uint32) ([]*projectdomain.ProjectDomain, error)
}

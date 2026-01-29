package postgres

import (
	"context"
	"database/sql"
	"errors"
	projectdomain "projectservice/internal/domain/project"
	posmapper "projectservice/internal/infrastructure/postgres/mapper"
	posmodels "projectservice/internal/infrastructure/postgres/models"
	"projectservice/internal/repository/storage"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (p *Postgres) FindByName(ctx context.Context, ownerId uint32, name string) (*projectdomain.ProjectDomain, error) {
	row := p.db.QueryRowContext(ctx, QuerieFindByName, ownerId, name)

	var pm posmodels.ProjectPosModel

	err := row.Scan(
		&pm.Id,
		&pm.OwnerId,
		&pm.Name,
		&pm.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, err
	}

	return posmapper.ModelToDomain(&pm), nil
}

func (p *Postgres) Save(ctx context.Context, proj *projectdomain.ProjectDomain) error {
	pm := posmapper.DomainToModel(proj)

	_, err := p.db.ExecContext(ctx, QuerieSave, pm.OwnerId, pm.Name)

	return err
}

package postgres

import (
	"context"
	"database/sql"
	projectdomain "projectservice/internal/domain/project"
	posmapper "projectservice/internal/infrastructure/postgres/mapper"
	posmodels "projectservice/internal/infrastructure/postgres/models"
	"projectservice/internal/repository/storage"

	"github.com/lib/pq"
)

var (
	invalidId uint32 = 0
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (p *Postgres) Save(ctx context.Context, proj *projectdomain.ProjectDomain) (uint32, error) {
	pm := posmapper.DomainToModel(proj)

	row := p.db.QueryRowContext(ctx, QuerieSave, pm.OwnerId, pm.Name)

	var id uint32

	err := row.Scan(
		&id,
	)

	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return invalidId, storage.ErrAlreadyExists
			}
		}
		return invalidId, err
	}

	return uint32(id), nil
}

func (p *Postgres) Delete(ctx context.Context, ownerId uint32, projectId uint32) error {
	res, err := p.db.ExecContext(ctx, QuerieDelete, projectId, ownerId)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if ra == 0 {
		return storage.ErrNotFound
	}

	return nil
}

func (p *Postgres) GetAll(ctx context.Context, ownerId uint32) ([]*projectdomain.ProjectDomain, error) {
	rows, err := p.db.QueryContext(ctx, QuerieGetAll, ownerId)
	if err != nil {
		return nil, err
	}

	var projects []*posmodels.ProjectPosModel
	for rows.Next() {
		project := &posmodels.ProjectPosModel{}

		err := rows.Scan(
			&project.Id,
			&project.OwnerId,
			&project.Name,
			&project.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(projects) == 0 {
		return nil, storage.ErrNotFound
	}

	return posmapper.ModelsToDomain(projects), nil
}

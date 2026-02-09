package postgres

import (
	"context"
	"database/sql"
	taskdomain "taskservice/internal/domain/task"
	posmapper "taskservice/internal/infrastructure/postgres/mapper"
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

func (p *Postgres) Save(ctx context.Context, td *taskdomain.TaskDomain) (uint32, error) {
	model := posmapper.TaskDomainToModel(td)

	row := p.db.QueryRowContext(ctx, QuerieCreate, model.ProjectId, model.Description, model.Deadline)

	var id uint32

	err := row.Scan(&id)
	if err != nil {
		return invalidId, err
	}

	return id, nil
}

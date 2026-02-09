package posmodels

import (
	"database/sql"
	"time"
)

type TaskPosModel struct {
	Id          uint32       `db:"id"`
	ProjectId   uint32       `db:"project_id"`
	Description string       `db:"description"`
	Deadline    sql.NullTime `db:"deadline"`
}

func NewTaskPosModel(id uint32, projectId uint32, description string, deadline time.Time) *TaskPosModel {
	dline := sql.NullTime{
		Time:  deadline,
		Valid: !deadline.IsZero(),
	}

	return &TaskPosModel{
		Id:          id,
		ProjectId:   projectId,
		Description: description,
		Deadline:    dline,
	}
}

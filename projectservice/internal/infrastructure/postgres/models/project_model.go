package posmodels

import "time"

type ProjectPosModel struct {
	Id        uint32    `db:"id"`
	OwnerId   uint32    `db:"owner_id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

func NewProjectPosModel(ownerId uint32, name string) *ProjectPosModel {
	return &ProjectPosModel{
		OwnerId: ownerId,
		Name:    name,
	}
}

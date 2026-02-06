package createdto

import "time"

type CreateRequest struct {
	ProjectId   uint32    `json:"project_id" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Deadline    time.Time `json:"deadline"`
}

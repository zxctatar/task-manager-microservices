package posmapper

import (
	taskdomain "taskservice/internal/domain/task"
	posmodels "taskservice/internal/infrastructure/postgres/models"
)

func TaskDomainToModel(td *taskdomain.TaskDomain) *posmodels.TaskPosModel {
	return posmodels.NewTaskPosModel(
		td.Id,
		td.ProjectId,
		td.Description,
		td.Deadline,
	)
}

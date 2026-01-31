package interfaces

import (
	"context"
	getallmodel "projectservice/internal/usecase/models/getallprojects"
)

type GetAllProjectsUsecase interface {
	Execute(ctx context.Context, in *getallmodel.GetAllProjectsInput) (*getallmodel.GetAllProjectsOutput, error)
}

package interfaces

import (
	"context"
	deletemodel "projectservice/internal/usecase/models/deleteproject"
)

type DeleteProjectUsecase interface {
	Execute(ctx context.Context, in *deletemodel.DeleteProjectInput) (*deletemodel.DeleteProjectOutput, error)
}

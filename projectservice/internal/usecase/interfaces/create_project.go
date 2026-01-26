package interfaces

import (
	createmodel "projectservice/internal/usecase/models/createproject"

	"github.com/gin-gonic/gin"
)

type CreateProjectUsecase interface {
	Execute(ctx *gin.Context, in *createmodel.CreateProjectInput) (*createmodel.CreateProjectOutput, error)
}
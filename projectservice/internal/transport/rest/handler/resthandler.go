package resthandler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type RestHandler struct {
	log *slog.Logger
}

func NewHandler(log *slog.Logger) *RestHandler {
	return &RestHandler{
		log: log,
	}
}

func (h *RestHandler) CreateProject(ctx *gin.Context) {
	panic("not implemented")
}

func (h *RestHandler) RemoveProject(ctx *gin.Context) {
	panic("not implemented")
}

func (h *RestHandler) GetAllProjects(ctx *gin.Context) {
	panic("not implemented")
}

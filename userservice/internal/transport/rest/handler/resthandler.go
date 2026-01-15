package resthandler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type RestHandler struct {
	log *slog.Logger
}

func NewRestHandler(log *slog.Logger) *RestHandler {
	return &RestHandler{
		log: log,
	}
}

func (h *RestHandler) Registration(ctx *gin.Context) {
	panic("not implemented")
}

func (h *RestHandler) Login(ctx *gin.Context) {
	panic("not implemented")
}

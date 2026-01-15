package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func timeoutResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusRequestTimeout, "timeout")
}

func TimeoutMiddleware(tout time.Duration) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(tout),
		timeout.WithResponse(timeoutResponse),
	)
}

package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSessionMiddleware(log *slog.Logger) gin.HandlerFunc {
	const op = "middleware.GetSessionMiddleware"
	return func(ctx *gin.Context) {
		sessionId, err := ctx.Cookie("sessionId")
		if err != nil {
			log.Info("a request arrived without a sessionId", slog.String("op", op))
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "needed cookie with sessionId",
			})
			ctx.Abort()
			return
		}
		ctx.Set("sessionId", sessionId)
		ctx.Next()
	}
}

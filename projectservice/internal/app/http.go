package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"projectservice/internal/config"
	"projectservice/internal/transport/rest"
	"projectservice/internal/transport/rest/middleware"

	"github.com/gin-gonic/gin"
)

func mustLoadHttpServer(cfg *config.Config, log *slog.Logger) *rest.RestServer {
	gin.SetMode(cfg.RestConf.Mode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.TimeoutMiddleware(cfg.RestConf.RequestTimeout))

	serv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.RestConf.Port),
		Handler:      router,
		ReadTimeout:  cfg.RestConf.ReadTimeout,
		WriteTimeout: cfg.RestConf.WriteTimeout,
	}

	return rest.NewRestServer(log, serv)
}

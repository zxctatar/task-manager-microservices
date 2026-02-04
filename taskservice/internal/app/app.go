package app

import (
	"context"
	"taskservice/internal/config"
	"taskservice/internal/infrastructure/grpc/userservice"
	"taskservice/internal/transport/rest"
	resthandler "taskservice/internal/transport/rest/handler"
	"taskservice/pkg/logger"
)

type App struct {
	cfg        *config.Config
	restServer *rest.RestServer
	client     *userservice.UserServiceClient
}

func NewApp() *App {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.LoggerConf.Level)

	client := userservice.NewUserServiceClient(log, cfg.ConnectionsConf.UserServConnConf.Host, cfg.ConnectionsConf.UserServConnConf.Port)
	handl := resthandler.NewRestHandler(log)

	restServer := mustLoadRestServer(cfg, log, handl, client)

	return &App{
		cfg:        cfg,
		restServer: restServer,
		client:     client,
	}
}

func (a *App) Run() {
	a.restServer.MustStart()
}

func (a *App) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RestConf.ShutdownTimeout)
	defer cancel()

	a.restServer.Stop(ctx)
	a.client.Stop()
}

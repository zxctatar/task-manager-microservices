package app

import (
	"context"
	"database/sql"
	"taskservice/internal/config"
	"taskservice/internal/infrastructure/grpc/userservice"
	"taskservice/internal/infrastructure/postgres"
	"taskservice/internal/transport/rest"
	resthandler "taskservice/internal/transport/rest/handler"
	createuc "taskservice/internal/usecase/implementations/createtask"
	"taskservice/pkg/logger"
)

type App struct {
	cfg        *config.Config
	restServer *rest.RestServer
	client     *userservice.UserServiceClient
	db         *sql.DB
}

func NewApp() *App {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.LoggerConf.Level)

	db := mustLoadPostgres(cfg)

	postgres := postgres.NewPostgres(db)

	createUC := createuc.NewCreateTaskUC(log, postgres)

	client := userservice.NewUserServiceClient(log, cfg.ConnectionsConf.UserServConnConf.Host, cfg.ConnectionsConf.UserServConnConf.Port)
	handl := resthandler.NewRestHandler(log, createUC)

	restServer := mustLoadRestServer(cfg, log, handl, client)

	return &App{
		cfg:        cfg,
		restServer: restServer,
		client:     client,
		db:         db,
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
	a.db.Close()
}

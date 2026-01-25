package app

import (
	"context"
	"log/slog"
	"projectservice/internal/config"
	userserviceclient "projectservice/internal/infrastructure/grpc/userservice"
	"projectservice/internal/transport/rest"
	"projectservice/pkg/logger"
)

type App struct {
	log  *slog.Logger
	cfg  *config.Config
	serv *rest.RestServer
}

func NewApp() *App {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.LoggerConf.Level)

	sessionValid := userserviceclient.NewUserServiceClient(log, cfg.ConnectionsConf.UserServConnConf.Host, cfg.ConnectionsConf.UserServConnConf.Port)

	serv := mustLoadHttpServer(cfg, log, sessionValid)

	return &App{
		log:  log,
		cfg:  cfg,
		serv: serv,
	}
}

func (a *App) Run() {
	a.serv.MustStart()
}

func (a *App) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RestConf.ShutdownTimeout)
	defer cancel()
	a.serv.Stop(ctx)
}

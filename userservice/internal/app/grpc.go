package app

import (
	"log/slog"
	"userservice/internal/config"
	grpcserv "userservice/internal/transport/grpc"
	userservicev1 "userservice/proto/userservice"
)

func mustLoadGRPCServer(cfg *config.Config, log *slog.Logger, handl userservicev1.UserServiceServer) *grpcserv.GRPCServer {
	return grpcserv.NewGRPCServer(log, cfg.GrpcConf.Port, handl)
}

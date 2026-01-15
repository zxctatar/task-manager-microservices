package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"userservice/internal/config"
	"userservice/internal/transport/rest"
	"userservice/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.MustLoad()
	log := logger.SetupLogger(config.LogConf.Level)

	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Recovery())

	serv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.RestConf.Port),
		Handler:      router,
		WriteTimeout: config.RestConf.WriteTimeout,
		ReadTimeout:  config.RestConf.ReadTimeout,
	}

	restServer := rest.NewRestServer(log, serv)

	go restServer.MustStart()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)

	<-quitCh

	ctx, cancel := context.WithTimeout(context.Background(), config.RestConf.ShutdownTimeout*time.Second)
	defer cancel()

	restServer.Stop(ctx)
}

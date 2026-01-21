package main

import (
	"os"
	"os/signal"
	"projectservice/internal/app"
	"syscall"
)

func main() {
	app := app.NewApp()

	go app.Run()

	sysChan := make(chan os.Signal, 1)
	signal.Notify(sysChan, syscall.SIGINT, syscall.SIGTERM)

	<-sysChan

	app.Stop()
}

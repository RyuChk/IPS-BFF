package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ZecretBone/ips-bff/cmd/bff-grpc/di"
	"github.com/ZecretBone/ips-bff/internal/config"
)

func main() {
	config.LoadConfig()

	appName := "admin-grpc"
	if err := os.Setenv("APP_NAME", appName); err != nil {
		log.Panic("error while setting APP_NAME")
	}

	_, cleanUpFunc, err := di.InitializeContainer()
	if err != nil {
		log.Panic("Error While building bff-service grpc server")
	}
	defer cleanUpFunc()

	closeCh := make(chan os.Signal, 1)
	signal.Notify(closeCh, os.Interrupt, syscall.SIGTERM)

	<-closeCh
}

package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ZecretBone/ips-bff/cmd/bff-api/di"
	"github.com/ZecretBone/ips-bff/internal/config"
	"github.com/rs/zerolog/log"
)

func main() {
	config.LoadConfig()

	if err := os.Setenv("APP_NAME", "bff-api"); err != nil {
		panic(err)
	}

	_, cleanup, err := di.InitializeContainer()
	if err != nil {
		log.Panic().Err(err).Msg("error while inject dependency wire")
		panic(err)
	}
	defer cleanup()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
}

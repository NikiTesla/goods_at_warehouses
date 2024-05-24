package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/NikiTesla/goods_at_warehouses/pkg/environment"
	"github.com/NikiTesla/goods_at_warehouses/pkg/jsonrpc"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.WithError(err).Fatal("can't load env variables")
	}
	configFile := os.Getenv("CONFIG_FILE")
	env, err := environment.NewEnvironment(configFile)
	if err != nil {
		log.WithError(err).Fatal("can't load environment")
	}

	server := jsonrpc.NewServer(env)
	if err = server.Run(env.Config.Port); err != nil {
		log.WithError(err).Fatal("error occured while running server")
	}
}

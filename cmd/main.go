package main

import (
	"log"
	"os"

	"github.com/NikiTesla/goods_at_warehouses/pkg/environment"
	"github.com/NikiTesla/goods_at_warehouses/pkg/jsonrpc"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("can't load env variables, err:", err)
	}

	configFile := os.Getenv("CONFIGFILE")
	env, err := environment.NewEnvironment(configFile)
	if err != nil {
		log.Fatal("can't load environment, err:", err)
	}

	server := jsonrpc.NewServer(env)

	if err = server.Run(env.Config.Port); err != nil {
		log.Fatal("error occured while running server, error:", err)
	}
}

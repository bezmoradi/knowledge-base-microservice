package main

import (
	"github.com/bezmoradi/knowledge-base-microservice/src/config"
	"github.com/bezmoradi/knowledge-base-microservice/src/databases"
	"github.com/bezmoradi/knowledge-base-microservice/src/services"
)

func init() {
	config.LoadEnv()
}

func main() {
	if err := databases.ConnectToMongoDB(); err != nil {
		panic(err)
	}

	defer databases.DisconnectFromMongoDB()

	services.HandleCron()
}

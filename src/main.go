package main

import (
	"github.com/behzadmoradi/blog-microservices-cronjob-go/src/config"
	"github.com/behzadmoradi/blog-microservices-cronjob-go/src/databases"
	"github.com/behzadmoradi/blog-microservices-cronjob-go/src/services"
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

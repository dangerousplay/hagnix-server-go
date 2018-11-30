package main

import (
	"github.com/InVisionApp/go-logger"
	"github.com/kataras/iris"
	"hagnix-server-go1/config"
	"hagnix-server-go1/database"
	"hagnix-server-go1/routes"
	"hagnix-server-go1/service"
	"os"
)

var logger = log.NewSimple()

func main() {
	app := iris.New()
	app.Logger().SetLevel("disable")

	logger.Info("Starting ROTMG Server...")

	config.Init()
	database.Init()
	service.Init()

	routes.RegisterRoutes(app)

	app.Run(iris.Addr(":" + port()))
}

func port() string {
	port := os.Getenv("PORT")

	if len(port) < 1 {
		port = "80"
	}

	return port
}

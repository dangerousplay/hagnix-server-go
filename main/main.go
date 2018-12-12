package main

import (
	"github.com/ivahaev/go-logger"
	"github.com/kataras/iris"
	"hagnix-server-go1/config"
	"hagnix-server-go1/database"
	"hagnix-server-go1/routes"
	"hagnix-server-go1/service"
	"os"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("disable")

	logger.Info("Starting ROTMG Server...")

	config.Init()
	database.Init()
	service.Init()

	routes.RegisterRoutes(app)

	err := app.Run(iris.Addr("127.0.0.1:" + port()))

	if err != nil {
		logger.Error("error on start listening: ", err)
	}
}

func debug(ctx iris.Context) {
	logger.Info(ctx.Method() + " " + ctx.RequestPath(true))
	ctx.Next()
}

func port() string {
	port := os.Getenv("PORT")

	if len(port) < 1 {
		port = "80"
	}

	return port
}

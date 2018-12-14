package main

import (
	"github.com/ivahaev/go-logger"
	"github.com/kataras/iris"
	"hagnix-server-go1/config"
	"hagnix-server-go1/database"
	"hagnix-server-go1/redis"
	"hagnix-server-go1/routes"
	"hagnix-server-go1/service"
	"os"
	"strconv"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("disable")

	logger.Info("Starting ROTMG Server...")

	config.Init()
	database.Init()
	service.Init()
	redis.InitRedis()

	routes.RegisterRoutes(app)

	startWebServer(app)
}

func port() string {
	port := os.Getenv("PORT")

	if len(port) < 1 {
		port = "80"
	}

	return port
}

func startWebServer(application *iris.Application) {
	configs := iris.WithConfiguration(iris.Configuration{EnableOptimizations: true})

	TLS := os.Getenv("TLS")

	useTLS, _ := strconv.ParseBool(TLS)

	if !useTLS {
		err := application.Run(iris.Addr("127.0.0.1:"+port()), configs)

		if err != nil {
			logger.Error("error on start listening: ", err)
		}
	} else {
		domain := os.Getenv("DOMAIN")
		email := os.Getenv("TLS_EMAIL")

		err := application.Run(iris.AutoTLS("127.0.0.1:"+port(), domain, email), configs)

		if err != nil {
			logger.Error(err)
		}
	}
}

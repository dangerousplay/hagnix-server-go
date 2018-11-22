package main

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes"
	"os"
)

func main() {
	app := iris.New()

	database.Init()

	database.GetDBEngine().Sync(&models.Accounts{})
	database.GetDBEngine().Sync(&models.Death{})

	routes.RegisterRoutes(app)

	app.Run(iris.Addr(":" + port()))
}

func port() string {
	port := os.Getenv("PORT")

	if len(port) < 1 {
		port = "8080"
	}

	return port
}

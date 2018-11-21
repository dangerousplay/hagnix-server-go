package main

import (
	"github.com/kataras/iris"
	"hagnix-server-go/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if len(port) < 1 {
		port = ":8080"
	}

	app := iris.New()

	routes.RegisterRoutes(app)

	app.Run(iris.Addr(":" + port))
}

package main

import (
	"fmt"
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/routes"
	"os"
)

func main() {
	app := iris.New()

	database.Init()

	database.GetDBEngine().Iterate(&database.Death{}, func(idx int, bean interface{}) error {
		bean2, _ := bean.(database.Death)

		test := bean2.Name
		fmt.Printf(test)
		return nil
	})

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

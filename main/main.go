package main

import (
	"encoding/xml"
	"fmt"
	"github.com/InVisionApp/go-logger"
	"github.com/kataras/iris"
	"hagnix-server-go1/config"
	"hagnix-server-go1/database"
	"hagnix-server-go1/routes"
	"hagnix-server-go1/routes/account"
	"os"
)

var logger = log.NewSimple()

func main() {
	app := iris.New()
	app.Logger().SetLevel("disable")

	bytess, err := xml.MarshalIndent(&account.AccountXML{}, " ", "  ")

	if err != nil {
		panic(err)
	}

	str := string(bytess)

	fmt.Println(str)

	logger.Info("Starting ROTMG Server...")

	config.Init()
	database.Init()

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

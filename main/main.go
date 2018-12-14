package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
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

	byts, _ := base64.StdEncoding.DecodeString("AAAAGDQBAAABvgIAAABYAwABqBMEAAAAAAUAAAAABgAAAZsHAAABmwgAAAAXCQAAABcKAAAACwsAAAAADAAAABENAAAAAA4AAAAADwAAAAAQAAAAABEAAAAAEgAAAAATAAAAABQAAAB0FQAAAAAWAAAAABcAAAAAGAAAAAA=")

	buf := bytes.NewBuffer(byts)

	test, _ := buf.ReadByte()

	for {
		if test < 0 {
			break
		}

		var y int32
		err := binary.Read(buf, binary.BigEndian, &y)

		if err != nil {
			break
		}

		fmt.Printf("%d: %d\n", test, y)

		test, _ = buf.ReadByte()
	}

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

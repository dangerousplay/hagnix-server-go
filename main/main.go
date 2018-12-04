package main

import (
	"bufio"
	"github.com/InVisionApp/go-logger"
	"github.com/kataras/iris"
	"hagnix-server-go1/config"
	"hagnix-server-go1/database"
	"hagnix-server-go1/routes"
	"hagnix-server-go1/service"
	"net"
	"os"
	"strconv"
	"strings"
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

	app.UseGlobal(debug)

	err := app.Run(iris.Addr("127.0.0.1:" + port()))

	if err != nil {
		logger.Error("error on start listening: ", err)
	}
}

func debug(ctx iris.Context) {
	logger.Info(ctx.Method() + " " + ctx.RequestPath(true))
	ctx.Next()
}

func handleConnection(c net.Conn) {
	logger.Infof("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			logger.Info(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		result := strconv.Itoa(12) + "\n"
		c.Write([]byte(string(result)))
	}
	c.Close()
}

func port() string {
	port := os.Getenv("PORT")

	if len(port) < 1 {
		port = "80"
	}

	return port
}

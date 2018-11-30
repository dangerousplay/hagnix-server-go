package char

import (
	"github.com/InVisionApp/go-logger"
	"github.com/kataras/iris"
)

var logger = log.NewSimple()

func RegisterRoutes(app *iris.Application) {
	capp := app.Party("/char")

	capp.Get("/delete", deleteChar)
	capp.Get("/list", handleList)
}

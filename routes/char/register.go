package char

import (
	"github.com/kataras/iris"
)

func RegisterRoutes(app *iris.Application) {
	capp := app.Party("/char")

	capp.Post("/delete", deleteChar)
	capp.Post("/list", handleList)
	capp.Post("/purchaseClassUnlock", handlePurchase)
	capp.Post("/fame", handleFame)
}

package account

import "github.com/kataras/iris"

func RegisterAccountRoutes(app *iris.Application) {
	raccount := app.Party("/account")

	raccount.Get("/acceptTOS", handleAccepTOS)
}

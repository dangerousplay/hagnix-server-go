package account

import "github.com/kataras/iris"

func RegisterAccountRoutes(app *iris.Application) {
	raccount := app.Party("/account")

	raccount.Get("/acceptTOS", handleAccepTOS)
}

func validateInputLogin(ctx iris.Context) bool {
	guid := ctx.URLParam("guid")
	password := ctx.URLParam("password")

	return validateLogin(guid, password)
}

func validateLogin(guid string, password string) bool {

	if len(guid) > 0 && len(password) > 0 {
		return true
	} else {
		return false
	}
}

package appn

import "github.com/kataras/iris"

func RegisterRouters(app *iris.Application) {
	rapp := app.Party("/app")

	rapp.Post("/init", handleInit)
	rapp.Post("/getLanguageStrings", handleLanguage)
	rapp.Post("/globalNews", handleGlobalNews)
}

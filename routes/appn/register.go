package appn

import "github.com/kataras/iris"

func RegisterRouters(app *iris.Application) {
	rapp := app.Party("/app")

	rapp.Get("/init", handleInit)
	rapp.Get("/getLanguageStrings", handleLanguage)
}

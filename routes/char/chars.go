package chars

import "github.com/kataras/iris"

func RegisterRoutes(app *iris.Application) {
	app.Delete("/char", deleteChar)
}

func deleteChar(ctx iris.Context) {

}

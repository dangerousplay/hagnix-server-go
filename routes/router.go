package routes

import "github.com/kataras/iris"

func RegisterRoutes(app *iris.Application){
	app.Get("/crossdomain", HandleCrossDomain)
}

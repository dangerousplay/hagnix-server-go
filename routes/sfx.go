package routes

import (
	"github.com/kataras/iris"
)

func handleSfx(app *iris.Application, dir string) {
	app.StaticWeb("/sfx", dir+"/sfx")
}

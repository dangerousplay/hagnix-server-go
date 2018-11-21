package routes

import (
	"github.com/kataras/iris"
)

func HandleSfx(app *iris.Application, dir string) {
	app.StaticWeb("/sfx", dir+"/sfx")
}

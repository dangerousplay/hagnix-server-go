package routes

import (
	"github.com/kataras/iris"
)

func handleTexture(app *iris.Application, dir string) {
	app.StaticWeb("/texture", dir+"/texture")
}

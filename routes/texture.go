package routes

import (
	"github.com/kataras/iris"
)

func HandleTexture(app *iris.Application, dir string) {
	app.StaticWeb("/texture", dir+"/texture")
}

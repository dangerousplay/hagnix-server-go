package routes

import (
	"github.com/kataras/iris"
	"os"
	"path/filepath"
)

func RegisterRoutes(app *iris.Application) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		panic(err)
	}

	app.Get("/crossdomain", HandleCrossDomain)
	HandleSfx(app, dir)
	HandleTexture(app, dir)
}

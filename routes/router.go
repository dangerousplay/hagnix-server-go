package routes

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/routes/account"
	"hagnix-server-go1/routes/appn"
	"os"
	"path/filepath"
)

func RegisterRoutes(app *iris.Application) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		panic(err)
	}

	app.Get("/crossdomain", handleCrossDomain)
	app.Get("/fame/list", handleFameList)
	app.Get("/health", handleHealth)
	account.RegisterAccountRoutes(app)
	appn.RegisterRouters(app)
	handleSfx(app, dir)
	handleTexture(app, dir)
}

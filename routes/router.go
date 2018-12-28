package routes

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/routes/account"
	"hagnix-server-go1/routes/appn"
	"hagnix-server-go1/routes/char"
	"hagnix-server-go1/routes/mysterybox"
	"hagnix-server-go1/routes/package"
	"os"
	"path/filepath"
)

var dir string

func RegisterRoutes(app *iris.Application) {

	direct, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		panic(err)
	}

	dir = direct

	app.Get("/crossdomain.xml", handleCrossDomain)
	app.Get("/fame/list", handleFameList)
	app.Get("/health", handleHealth)
	app.Post("/picture/get", handleGetPicture)
	account.RegisterAccountRoutes(app)
	appn.RegisterRouters(app)
	char.RegisterRoutes(app)
	_package.RegisterRoutes(app)
	mysterybox.RegisterRoutes(app)
	handleSfx(app, dir)
	handleTexture(app, dir)
}

package mysterybox

import "github.com/kataras/iris"

func RegisterRoutes(app *iris.Application) {
	mapp := app.Party("mysterybox")

	mapp.Get("/getBoxes", handleGetBoxes)
}

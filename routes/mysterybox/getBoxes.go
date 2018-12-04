package mysterybox

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/routes/messages"
)

func handleGetBoxes(ctx iris.Context) {
	//TODO implement Mystery Box Get Boxes
	ctx.XML(messages.DefaultSuccess)
}

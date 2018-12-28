package routes

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/utils"
)

func handleGetPicture(ctx iris.Context) {
	id := ctx.PostValue("id")

	if len(id) < 1 {
		ctx.XML(messages.BadRequest)
		return
	}
	file := id + ".png"

	err := ctx.SendFile(dir+"/texture/"+file, file)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}
}

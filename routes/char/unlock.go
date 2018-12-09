package char

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/routes/messages"
)

func handlePurchase(ctx iris.Context) {
	guid := ctx.PostValue("guid")
	password := ctx.PostValue("password")
	class := ctx.PostValue("classType")

	//TODO implement ClassTypeID from list of IDs
	if len(guid) < 1 || len(password) < 1 || len(class) < 1 {
		ctx.XML(messages.BadRequest)
		return
	}
}

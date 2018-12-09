package char

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/utils"
)

func handleFame(ctx iris.Context) {
	accountId := ctx.PostValue("accountId")
	charId := ctx.PostValue("charId")

	var death models.Death

	success, err := database.GetDBEngine().Where("accId = ? AND charId = ?", accountId, charId).Get(&death)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if !success {
		ctx.XML(messages.BadRequest)
		return
	}

	//TODO implement Fame

}

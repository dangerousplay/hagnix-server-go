package char

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/utils"
	"hagnix-server-go1/service"
)

func handleFame(ctx iris.Context) {
	accountId := ctx.PostValue("accountId")
	charId := ctx.PostValue("charId")

	if len(accountId) < 1 || len(charId) < 1 {
		ctx.XML(messages.BadRequest)
		return
	}

	var death models.Death

	success, err := database.GetDBEngine().Where("accId = ? AND chrId = ?", accountId, charId).Get(&death)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	if !success {
		ctx.XML(messages.BadRequest)
		return
	}

	//TODO implement Fame

	acxml, acc, err := service.GetAccountService().VerifyGenerateAccountXMLbyId(accountId)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	if acxml == nil || acc == nil {
		return
	}

	var character models.Characters

	success, err = database.GetDBEngine().Where("accId = ?", accountId).Get(&character)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	if !success {
		ctx.XML(messages.DefaultError)
		return
	}

	deathXML, err := service.GetFameService().GetDeathFame(acc, &character, &death)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	ctx.XML(deathXML)
}

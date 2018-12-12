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

	var death models.Death

	success, err := database.GetDBEngine().Where("accId = ? AND charId = ?", accountId, charId).Get(&death)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	if !success {
		ctx.XML(messages.BadRequest)
		return
	}

	var account models.Accounts

	success, err = database.GetDBEngine().Id(accountId).Get(&account)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	if !success {
		ctx.XML(messages.DefaultError)
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

}

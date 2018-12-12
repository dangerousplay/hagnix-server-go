package char

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/utils"
	"hagnix-server-go1/service"
)

func deleteChar(ctx iris.Context) {
	guid := ctx.PostValue("guid")
	password := ctx.PostValue("password")
	charId := ctx.PostValue("charId")

	if len(guid) < 1 || len(password) < 1 || len(charId) < 1 {
		ctx.XML(messages.BadRequest)
		return
	}

	account, err := service.GetAccountService().Verify(guid, password)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	if account == nil {
		ctx.XML(messages.BadRequest)
		return
	}

	rows, err := database.GetDBEngine().Where("accId = ? AND charId = ?", account.Id, charId).Delete(&models.Characters{})

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	if rows > 0 {
		ctx.XML(messages.DefaultSuccess)
	} else {
		ctx.XML(messages.DefaultError)
	}

}

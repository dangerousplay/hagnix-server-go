package account

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/utils"
	"hagnix-server-go1/service"
	"regexp"
)

func handleSetName(ctx iris.Context) {
	guid := ctx.URLParam("guid")
	password := ctx.URLParam("password")
	name := ctx.URLParam("name")

	account, err := service.GetAccountService().Verify(guid, password)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if !regexp.MustCompile("^[a-zA-Z]+$").MatchString(name) || account.Namechosen != 0 {
		ctx.XML(messages.BadRequest)
		return
	}
	ex, err := service.GetAccountService().NameExists(name)

	if ex {
		ctx.XML(messages.Error{RawXml: "Name already exists"})
		return
	}

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	rows, err := database.GetDBEngine().Cols("name").Where("uuid = ?", account.Id).Update(&models.Accounts{Name: name, Namechosen: 1})

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if rows != 1 {
		ctx.XML(messages.DefaultError)
	} else {
		ctx.XML(messages.DefaultSuccess)
	}
}

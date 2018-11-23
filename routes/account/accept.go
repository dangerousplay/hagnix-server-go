package account

import (
	"github.com/InVisionApp/go-logger"
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/utils"
	"hagnix-server-go1/service"
)

var logger = log.NewSimple()

func handleAccepTOS(ctx iris.Context) {
	guid := ctx.URLParam("guid")
	password := ctx.URLParam("password")

	if len(guid) < 1 && len(password) < 1 {
		ctx.XML(messages.Error{RawXml: "Bad request"})
		return
	}

	account, err := service.GetAccountService().VerifyOnly(guid, password)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if !account {
		ctx.XML(messages.Error{RawXml: "Account not found"})
		return
	}

	_, err = database.GetDBEngine().Cols("acceptedNewTos").Where("uuid = ?", guid).Update(&models.Accounts{Acceptednewtos: 1})

	if !utils.DefaultErrorHandler(ctx, err, logger) {
		ctx.XML(messages.Sucess{Message: "OK"})
	}

}

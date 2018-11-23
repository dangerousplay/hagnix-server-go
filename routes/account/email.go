package account

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/config"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/utils"
	"hagnix-server-go1/service"
)

func handleChangeEmail(ctx iris.Context) {
	guid := ctx.URLParam("guid")
	password := ctx.URLParam("password")
	newGuid := ctx.URLParam("newGuid")

	if len(newGuid) < 1 {
		ctx.StatusCode(400)
		ctx.XML(messages.Error{RawXml: "Bad Request"})
		return
	}

	account, err := service.GetAccountService().Verify(guid, password)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if account.Verified != 0 || !config.Config.ServerConfig.VerifyEmail {
		ctx.XML(messages.Sucess{})
		return
	}

	//TODO implement Email

}

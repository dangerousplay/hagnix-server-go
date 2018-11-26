package account

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/config"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/utils"
	"hagnix-server-go1/service"
)

const verifySucess = `<html>
    <body bgcolor="#000000">
        <div align="center">
            <font color="#FFFFFF">Your account is now verified.</font>
        </div>
    </body>
</html>`

const verifyFail = `<html>
<body bgcolor="#000000">
    <div align="center">
        <font color="#FFFFF">Could not verify your account. Please try again.</font>
    </div>
</body>
</html>`

func handleChangeEmail(ctx iris.Context) {
	guid := ctx.URLParam("guid")
	password := ctx.URLParam("password")
	newGuid := ctx.URLParam("newGuid")

	if len(newGuid) < 1 {
		ctx.StatusCode(400)
		ctx.XML(messages.BadRequest)
		return
	}

	account, err := service.GetAccountService().Verify(guid, password)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if account.Verified != 0 || !config.GetConfig().ServerConfig.VerifyEmail {
		ctx.XML(messages.DefaultSuccess)
		return
	}

	//TODO implement Email

}

func handleValidateEmail(ctx iris.Context) {
	authToken := ctx.URLParam("authToken")

	if len(authToken) < 1 {
		ctx.XML(messages.BadRequest)
		return
	}

	rows, err := database.GetDBEngine().Cols("verified").Where("authToken = ?", authToken).Update(&models.Accounts{Verified: 1})

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if rows != 1 {
		ctx.HTML(verifyFail)
	} else {
		ctx.HTML(verifySucess)
	}

}

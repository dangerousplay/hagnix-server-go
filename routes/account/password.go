package account

import (
	"fmt"
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/utils"
	"hagnix-server-go1/service"
)

func handleChangePassword(ctx iris.Context) {
	guid := ctx.URLParam("guid")
	password := ctx.URLParam("password")
	newPassword := ctx.URLParam("newPassword")

	if validateLogin(guid, password) && len(password) < 1 {
		ctx.XML(messages.BadRequest)
		return
	}

	account, err := service.GetAccountService().Verify(guid, password)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	rows, err := database.GetDBEngine().Cols("password").Where("uuid = ?", account.Uuid).Update(&models.Accounts{Password: utils.HashStringSHA1(newPassword)})

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if rows > 0 {
		ctx.XML(messages.DefaultSuccess)
	} else {
		ctx.XML(messages.DefaultError)
	}
}

func handleForgotPassword(ctx iris.Context) {
	//TODO implement forgot password
}

func handleResetPassword(ctx iris.Context) {
	authToken := ctx.URLParam("authToken")
	newPassword := utils.RandomString(10)

	if len(authToken) < 1 {
		ctx.XML(messages.BadRequest)
		return
	}

	rows, err := database.GetDBEngine().Cols("password").Where("authToken = ?", authToken).Update(&models.Accounts{Password: utils.HashStringSHA1(newPassword)})

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if rows != 1 {
		ctx.XML(messages.DefaultError)
		return
	}

	ctx.HTML(
		fmt.Sprintf(`<html>
<body bgcolor=""#000000"">
    <div align=""center"">
        <font color="#FFFFFF">Your new password is %s, please note that passwords are CaSeSensItivE.</font>
    </div>
</body>
</html>`, newPassword))
}

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

var invalidEmail = &messages.Error{RawXml: "WebRegister.invalid_email_address"}
var alreadyUsed = &messages.Error{RawXml: "Error.emailAlreadyUsed"}
var emailError = &messages.Error{RawXml: "WebForgotPasswordDialog.emailError"}

func handleRegister(ctx iris.Context) {
	ignore := ctx.URLParam("ignore")
	entrytag := ctx.URLParam("entrytag")
	isAgeVerified := ctx.URLParam("isAgeVerified")
	newGuid := ctx.URLParam("newGuid")
	guid := ctx.URLParam("guid")
	newPassword := ctx.URLParam("newPassword")

	if len(ignore) < 1 || len(entrytag) < 1 || len(isAgeVerified) < 1 {
		ctx.XML(invalidEmail)
		return
	}

	if !regexp.MustCompile(`^([a-zA-Z0-9_\-\.]+)@([a-zA-Z0-9_\-\.]+)\.([a-zA-Z]{2,5})$`).MatchString(newGuid) {
		ctx.XML(emailError)
		return
	}

	exist, err := service.GetAccountService().Verify(guid, "")

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	if exist != nil && exist.Guest == 1 {
		exist2, err2 := service.GetAccountService().AccountExists(newGuid)

		if utils.DefaultErrorHandler(ctx, err2) {
			return
		}

		if exist2 {
			ctx.XML(alreadyUsed)
			return
		}

		rows, err := database.GetDBEngine().Where("uuid = ?", guid).Update(&models.Accounts{Name: newGuid, Uuid: newGuid, Guest: 0, Password: utils.HashStringSHA1(newPassword)})

		if utils.DefaultErrorHandler(ctx, err) {
			return
		}

		if rows != 1 {
			ctx.XML(messages.DefaultError)
			return
		}

		ctx.XML(messages.DefaultSuccess)

		return
	} else {
		_, err := service.GetAccountService().Register(newGuid, newPassword)

		if err != nil {
			ctx.XML(messages.DefaultError)
		} else {
			ctx.XML(messages.DefaultSuccess)
		}
	}

}

func handleVerify(ctx iris.Context) {
	guid := ctx.PostValue("guid")
	password := ctx.PostValue("password")

	if len(guid) < 1 || len(password) < 1 {
		ctx.XML(messages.BadRequest)
		return
	}

	account, err := service.GetAccountService().GenerateAccountXML(guid, password)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	ctx.XML(account)
}

func verifyAge(ctx iris.Context) {
	//TODO implement verify Age
}

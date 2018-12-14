package account

import "github.com/kataras/iris"

func RegisterAccountRoutes(app *iris.Application) {
	raccount := app.Party("/account")

	raccount.Post("/acceptTOS", handleAccepTOS)
	raccount.Post("/changeEmail", handleChangeEmail)
	raccount.Post("/changePassword", handleChangePassword)
	raccount.Post("/checkGiftCode", handleGiftCode)
	raccount.Post("/purchaseCharSlot", handlePurchaseCharSlot)
	raccount.Post("/purchasePackage", handlePurchasePackage)
	raccount.Post("/purchaseSkin", handlePurchaseSkin)
	raccount.Post("/register", handleRegister)
	raccount.Post("/resetPassword", handleResetPassword)
	raccount.Post("/setName", handleSetName)
	raccount.Post("/validateEmail", handleValidateEmail)
	raccount.Post("/verify", handleVerify)

}

func validateInputLogin(ctx iris.Context) bool {
	guid := ctx.URLParam("guid")
	password := ctx.URLParam("password")

	return validateLogin(guid, password)
}

func validateLogin(guid string, password string) bool {

	if len(guid) > 0 && len(password) > 0 {
		return true
	} else {
		return false
	}
}

package account

import "github.com/kataras/iris"

func RegisterAccountRoutes(app *iris.Application) {
	raccount := app.Party("/account")

	raccount.Get("/acceptTOS", handleAccepTOS)
	raccount.Get("/changeEmail", handleChangeEmail)
	raccount.Get("/changePassword", handleChangePassword)
	raccount.Get("/checkGiftCode", handleGiftCode)
	raccount.Get("/purchaseCharSlot", handlePurchaseCharSlot)
	raccount.Get("/purchasePackage", handlePurchasePackage)
	raccount.Get("/purchaseSkin", handlePurchaseSkin)
	raccount.Get("/register", handleRegister)
	raccount.Get("/resetPassword", handleResetPassword)
	raccount.Get("/setName", handleSetName)
	raccount.Get("/validateEmail", handleValidateEmail)
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

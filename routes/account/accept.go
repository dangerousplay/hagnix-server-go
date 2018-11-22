package account

import (
	"fmt"
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/service"
)

func handleAccepTOS(ctx iris.Context) {
	guid := ctx.URLParam("guid")
	password := ctx.URLParam("password")

	if len(guid) < 1 && len(password) < 1 {
		ctx.XML(messages.Error{RawXml: "Bad request"})
		return
	}

	account, err := service.GetAccountService().Verify(guid, password)

	if err != nil {
		ctx.StatusCode(500)
		ctx.XML(messages.Error{RawXml: "Something failed: " + err.Error()})
		fmt.Println(err)
		return
	}

	if account == nil {
		ctx.XML(messages.Error{RawXml: "Account not found"})
		return
	}

	if account.Acceptednewtos == 0 {
		ctx.XML(messages.Error{RawXml: "TOS already accepted!"})
		return
	} else {
		account.Acceptednewtos = 1
		_, err := database.GetDBEngine().Cols("acceptedNewTos").Where("uuid = ?", account.Uuid).Update(&account)

		if err != nil {
			fmt.Println(err)
		}

		ctx.XML(messages.Sucess{Message: "OK"})
	}

}

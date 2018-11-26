package account

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/utils"
	"hagnix-server-go1/service"
)

func handlePurchaseCharSlot(ctx iris.Context) {
	guid := ctx.URLParam("guid")
	password := ctx.URLParam("password")
	account, err := service.GetAccountService().Verify(guid, password)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if account == nil {
		ctx.XML(messages.BadRequest)
		return
	}

	stats := &models.Stats{}

	success, err := database.GetDBEngine().Cols("credits").Where("accId = ?", account.Id).Get(&stats)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if !success {
		ctx.XML(messages.DefaultError)
		return
	}

	nextCharPrice := nextCharSlotPrice(account)

	if stats.Credits < nextCharPrice {
		ctx.XML(messages.Error{RawXml: "Not enough Gold"})
		return
	}

	session := database.GetDBEngine().NewSession()

	defer session.Close()

	rows, err := session.Cols("credits").Where("accId = ?", account.Id).Update(&models.Stats{Credits: stats.Credits - nextCharPrice})

	if utils.DefaultErrorHandler(ctx, err, logger) {
		session.Rollback()
		return
	}

	if rows > 0 {
		rows, err = session.Cols("maxCharSlot").Where("id = ?", account.Id).Update(&models.Accounts{Maxcharslot: account.Maxcharslot + 1})

		if utils.DefaultErrorHandler(ctx, err, logger) || rows < 1 {
			session.Rollback()
		}

		err := session.Commit()

		if utils.DefaultErrorHandler(ctx, err, logger) {
			return
		}

		ctx.XML(messages.DefaultSuccess)
	} else {
		ctx.XML(messages.DefaultError)
		return
	}
}

func nextCharSlotPrice(account *models.Accounts) int {
	var price int

	if account.Maxcharslot == 1 {
		price = 600
	} else if account.Maxcharslot == 2 {
		price = 800
	} else {
		price = 1000
	}

	return price
}

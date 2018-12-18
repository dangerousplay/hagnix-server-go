package char

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/modelxml"
	"hagnix-server-go1/routes/utils"
	"hagnix-server-go1/service"
)

func handlePurchase(ctx iris.Context) {
	guid := ctx.PostValue("guid")
	password := ctx.PostValue("password")
	class := ctx.PostValue("classType")

	//TODO implement ClassTypeID from list of IDs
	if len(guid) < 1 || len(password) < 1 || len(class) < 1 {
		ctx.XML(messages.BadRequest)
		return
	}

	account, err := service.GetAccountService().VerifyOnly(guid, password)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	if !account {
		ctx.XML(messages.BadRequest)
		return
	}

	classXML := modelxml.GetItem(class)

	if classXML == nil {
		ctx.XML(messages.BadRequest)
		return
	}

	id, err := service.GetAccountService().VerifyAndGetId(guid, password)

	if id == nil {
		ctx.XML(messages.BadRequest)
		return
	}

	var stats *models.Stats

	_, err = database.GetDBEngine().Cols("credits").Where("accId = ?", id).Get(stats)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	if classXML.UnlockCost < stats.Credits {
		return
	}

	session := database.GetDBEngine().NewSession()

	defer session.Close()

	rows, err := session.Where("accId = ?", id).Update(&models.Stats{Credits: stats.Credits - classXML.UnlockCost})

	if utils.HandleSessionRowsUpdated(ctx, session, err, rows, 1) {
		return
	}

	session.Commit()

	ctx.XML(messages.DefaultSuccess)

}

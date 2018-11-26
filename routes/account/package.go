package account

import (
	"encoding/json"
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/utils"
	"hagnix-server-go1/service"
)

type packageContents struct {
	Items       []int `json:"items"`
	VaultChests int   `json:"vaultChests"`
	CharSlots   int   `json:"charSlots"`
}

func handlePurchasePackage(ctx iris.Context) {
	guid := ctx.URLParam("guid")
	password := ctx.URLParam("password")
	packageId := ctx.URLParam("packageId")

	if !validateLogin(guid, password) {
		ctx.XML(messages.BadRequest)
		return
	}

	account, err := service.GetAccountService().Verify(guid, password)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if account == nil {
		ctx.XML(messages.Error{RawXml: "Account not found"})
		return
	}

	packages := &models.Packages{}

	success, err := database.GetDBEngine().Where("id = ? AND endDate >= now()", packageId).Get(&packages)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if !success {
		ctx.XML(messages.Error{RawXml: "This package is not available any more"})
		return
	}

	stats := &models.Stats{}

	success, err = database.GetDBEngine().Cols("credits").Where("accId = ?", account.Id).Get(&stats)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if !success {
		ctx.XML(messages.DefaultError)
	}

	if stats.Credits < packages.Price {
		ctx.XML(messages.Error{RawXml: "Not enough Gold"})
		return
	}

	contents := &packageContents{}

	err = json.Unmarshal([]byte(packages.Contents), contents)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	newGift := utils.AppendCommaSpaceSeparated(account.Gifts, contents.Items)

	session := database.GetDBEngine().NewSession()

	defer session.Close()

	rows, err := session.Cols("gifts").Where("uuid = ?", account.Uuid).Update(&models.Accounts{Gifts: newGift})

	if utils.HandleSessionRowsUpdated(ctx, session, err, logger, rows, 1) {
		return
	}

	if contents.CharSlots > 0 {
		rows, err = session.Cols("maxCharSlot").Where("uuid = ?", account.Uuid).Update(&models.Accounts{Maxcharslot: account.Maxcharslot + contents.CharSlots})

		if utils.HandleSessionRowsUpdated(ctx, session, err, logger, rows, int64(contents.CharSlots)) {
			return
		}
	}

	if contents.VaultChests > 0 {
		rows, err := service.GetAccountService().CreateChests(account, contents.VaultChests)

		if utils.HandleSessionRowsUpdated(ctx, session, err, logger, rows, int64(contents.VaultChests)) {
			return
		}
	}

	session.Commit()

	ctx.XML(messages.Success{})

}

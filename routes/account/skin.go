package account

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/utils"
	"hagnix-server-go1/service"
	"strconv"
)

type skin struct {
	Type        string
	Purchasable bool
	Expires     bool
	Price       int
}

type listSkin []skin

var skins = listSkin{
	skin{Type: "900", Purchasable: false, Expires: false, Price: 90000},
	skin{Type: "902", Purchasable: false, Expires: false, Price: 90000},
	skin{Type: "834", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "835", Purchasable: true, Expires: false, Price: 600},
	skin{Type: "836", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "837", Purchasable: true, Expires: false, Price: 600},
	skin{Type: "838", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "839", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "840", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "841", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "842", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "843", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "844", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "845", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "846", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "847", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "848", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "849", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "850", Purchasable: false, Expires: true, Price: 900},
	skin{Type: "851", Purchasable: false, Expires: true, Price: 900},
	skin{Type: "852", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "853", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "854", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "855", Purchasable: true, Expires: false, Price: 900},
	skin{Type: "856", Purchasable: false, Expires: false, Price: 90000},
	skin{Type: "883", Purchasable: false, Expires: false, Price: 90000},
}

func (skins *listSkin) findSkin(typ string) *skin {
	var temp *skin = nil

	for _, v := range *skins {
		if v.Type == typ {
			temp = &v
		}
	}

	return temp
}

func handlePurchaseSkin(ctx iris.Context) {
	guid := ctx.URLParam("guid")
	password := ctx.URLParam("password")
	skinType := ctx.URLParam("skinType")

	account, err := service.GetAccountService().Verify(guid, password)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if account == nil || len(skinType) < 1 {
		ctx.XML(messages.BadRequest)
		return
	}

	skin := skins.findSkin(skinType)

	skinNumber, err := strconv.Atoi(skinType)

	stats := &models.Stats{}

	_, err2 := database.GetDBEngine().Cols("credits").Where("accId = ?", account.Id).Get(stats)

	if utils.DefaultErrorHandler(ctx, err2, logger) {
		return
	}

	if skin == nil || err != nil || stats.Credits < skin.Price {
		ctx.XML(messages.BadRequest)
		return
	}

	contains, pSkins, err := service.GetAccountService().ContainsAndGetSkin(account, skinNumber)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	if contains {
		ctx.XML(messages.BadRequest)
		return
	}

	pSkins = append(pSkins, skinNumber)

	session := database.GetDBEngine().NewSession()

	defer session.Close()

	rows, err := session.Cols("ownedSkins").Where("uuid = ?", account.Id).Update(&models.Accounts{Ownedskins: utils.ToCommaSpaceSeparated(pSkins)})

	if utils.HandleSessionRowsUpdated(ctx, session, err, logger, rows, 1) {
		return
	}

	rows, err = session.Cols("credits").Where("accId = ?").Update(&models.Stats{Credits: stats.Credits - skin.Price})

	if utils.HandleSessionRowsUpdated(ctx, session, err, logger, rows, 1) {
		return
	}

	ctx.XML(messages.DefaultSuccess)

	session.Commit()
}

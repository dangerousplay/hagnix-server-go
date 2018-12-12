package char

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/modelxml"
	"hagnix-server-go1/routes/utils"
	"hagnix-server-go1/service"
)

func handleList(ctx iris.Context) {
	guid := ctx.PostValue("guid")
	password := ctx.PostValue("password")

	if len(guid) < 1 {
		ctx.XML(messages.BadRequest)
		return
	}

	news, err := service.GetNewsService().GetNews()

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	servers := service.GetServerService().GetServers()

	if len(password) < 1 {
		chars := &modelxml.CharsXML{
			Account: modelxml.AccountXML{
				Name:              service.GetAccountService().GetRandomName(),
				Id:                0,
				Admin:             false,
				Banned:            false,
				PetYardType:       1,
				NextCharSlotPrice: service.NextCharSlotPriceByChars(1),
			},
			NextCharId:        1,
			MaxNumChars:       2,
			NewsXML:           news,
			Servers:           servers,
			Classes:           modelxml.Classes,
			MaxClassLevelList: modelxml.MaxClassLevels,
			ItemCosts:         modelxml.Items,
		}

		ctx.XML(chars)
		return
	}

	accountXML, account, err := service.GetAccountService().VerifyGenerateAccountXML(guid, password)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	if accountXML == nil {
		ctx.XML(messages.BadRequest)
		return
	}

	charId, err := service.GetAccountService().NextCharId(account)
	characters, err2 := service.GetAccountService().GetCharsXML(account)

	if utils.DefaultErrorHandler(ctx, err) || utils.DefaultErrorHandler(ctx, err2) {
		return
	}

	var tos *int

	if account.Acceptednewtos == 1 {
		tos = &account.Acceptednewtos
	}

	classes, err := service.GetAccountService().GetAvailableClasses(account)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	chars := &modelxml.CharsXML{
		Account:           *accountXML,
		NextCharId:        charId,
		MaxNumChars:       account.Maxcharslot,
		OwnedSkins:        account.Ownedskins,
		NewsXML:           news,
		Servers:           servers,
		TOSPopup:          tos,
		Char:              characters,
		Classes:           classes,
		MaxClassLevelList: modelxml.MaxClassLevels,
		ItemCosts:         modelxml.Items,
	}

	ctx.XML(chars)

}

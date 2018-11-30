package char

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/modelxml"
	"hagnix-server-go1/routes/utils"
	"hagnix-server-go1/service"
)

func handleList(ctx iris.Context) {
	guid := ctx.URLParam("guid")
	password := ctx.URLParam("password")

	if len(guid) < 1 || len(password) < 1 {
		ctx.XML(messages.BadRequest)
		return
	}

	accountXML, account, err := service.GetAccountService().VerifyGenerateAccountXML(guid, password)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	news, err := service.GetNewsService().GetNews()

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	servers := service.GetServerService().GetServers()

	if accountXML == nil {
		chars := &modelxml.CharsXML{
			Account: modelxml.AccountXML{
				Name:        "AST",
				Id:          0,
				Admin:       false,
				Banned:      false,
				PetYardType: 1,
			},
			NewsXML: news,
			Servers: modelxml.ServersWrapper{Servers: servers},
		}

		ctx.XML(chars)
		return
	} else {
		charId, err := service.GetAccountService().NextCharId(account)
		characters, err2 := service.GetAccountService().GetCharsXML(account)

		if utils.DefaultErrorHandler(ctx, err, logger) || utils.DefaultErrorHandler(ctx, err2, logger) {
			return
		}

		chars := &modelxml.CharsXML{
			Account:           *accountXML,
			NextCharId:        charId,
			MaxNumChars:       account.Maxcharslot,
			OwnedSkins:        account.Ownedskins,
			NewsXML:           news,
			Servers:           modelxml.ServersWrapper{Servers: servers},
			TOSPopup:          account.Acceptednewtos,
			Char:              characters,
			Classes:           modelxml.ClassWrapper{Classes: modelxml.Classes},
			MaxClassLevelList: modelxml.MaxClassWrapper{MaxClasses: modelxml.MaxClassLevels},
			ItemCosts:         modelxml.ItemsWrapper{ItemCost: modelxml.Items},
		}

		ctx.XML(chars)
	}

}

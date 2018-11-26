package account

import (
	"encoding/xml"
	"fmt"
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
var emailError = &messages.Error{RawXml:"WebForgotPasswordDialog.emailError"}

var dailyQuestDescriptions = []string {
	"This is the first quest of the day! Bring me a %s and I will reward you with a fortune token! But if you can complete all the quests, there will be an added bonus for you!",
	"Ahh, you have moved on to the second quest! If you bring me a %s I can pull out the magic bits and make another Fortune Token! If you finish my next quest, I will up the ante a bit...",
	"You again! Excellent Since you have been so helpful, I will use some specific parts arround here and make you TWO Fortune Tokens. All I need is a %s",
	"You again! Excellent Since you have been so helpful, I will use some specific parts arround here and make you TWO Fortune Tokens. All I need is a %s",
}

type vaultXML struct {
	XMLName xml.Name `xml:"Vault"`
	Chests []string  `xml:"Chest"`
}

type guildXML struct {
	XMLName xml.Name `xml:"Guild"`
	Id      int      `xml:"id,attr"`
	Rank    string   `xml:"Rank"`
	Name    string   `xml:"Name"`
	Fame    int      `xml:"Fame"`

}

type giftsXML struct {
	XMLName xml.Name `xml:"Gifts"`
}

type dailyQuestXML struct {
	XMLName xml.Name   `xml:"DailyQuest"`
	Description string `xml:"Description"`
	Image       string `xml:"Image"`
	Tier        int    `xml:"tier,attr"`
	Goal        int    `xml:"goal,attr"`
}

type statsXML struct {
	XMLName xml.Name `xml:"Stats"`
	ClassStats []classStatsXML
}

type classStatsXML struct {
	XMLName xml.Name `xml:"ClassStats"`
    ObjectType string `xml:"objectType,attr"`
	BestLevel int `xml:"BestLevel"`
	BestFame int `xml:"BestFame"`
}

type AccountXML struct {
	XMLName xml.Name `xml:"Account"`
	Id             int64     `xml:"AccountId"`
	Name           string    `xml:"Name"`
	Namechosen     int       `xml:"NameChosen"`
	Admin          bool      `xml:"Admin"`
	Verified       bool      `xml:"VerifiedEmail"`
	Credits        int       `xml:"Credits"`
	FortuneTokens  int       `xml:"FortuneTokens"`
	NextCharSlotPrice int    `xml:"NextCharSlotPrice"`
	BeginnerPackageTimeLeft int `xml:"BeginnerPackageTimeLeft"`
	PetYardType    int       `xml:"PetYardType"`
	ArenaTickets   int       `xml:"ArenaTickets"`
	IsAgeVerified  int       `xml:"IsAgeVerified"`
	Stats          statsXML
	DailyQuest     dailyQuestXML
	Guild          guildXML
	Gifts          giftsXML
	Vault          vaultXML

}

func handleRegister(ctx iris.Context){
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

	if utils.DefaultErrorHandler(ctx, err, logger){
		return
	}

	if exist != nil && exist.Guest == 1 {
		exist2, err2 := service.GetAccountService().AccountExists(newGuid)

		if utils.DefaultErrorHandler(ctx, err2, logger) {
			return
		}

		if exist2 {
			ctx.XML(alreadyUsed)
			return
		}

		rows, err := database.GetDBEngine().Where("uuid = ?", guid).Update(&models.Accounts{Name: newGuid, Uuid: newGuid, Guest: 0, Password: utils.HashStringSHA1(newPassword)})

		if utils.DefaultErrorHandler(ctx,err,logger) {
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

func handleVerify(ctx iris.Context){
	guid := ctx.URLParam("guid")
	password := ctx.URLParam("password")

	if len(guid) < 1 || len(password) < 1 {
		ctx.XML(messages.BadRequest)
		return
	}

	account, err := service.GetAccountService().Verify(guid, password)

	if utils.DefaultErrorHandler(ctx, err, logger) {
		return
	}

	stats := &models.Stats{}

	success, err := database.GetDBEngine().Where("accId = ?", account.Id).Get(stats)

	if utils.DefaultErrorHandler(ctx, err, logger){
		return
	}

	if !success {
		ctx.XML(messages.DefaultError)
		return
	}

	var verifiedEmail = false

	if account.Verified == 1 {
		verifiedEmail = true
	}

	var admin = false

	if account.Rank > 2 {
		admin = true
	}

	var classes []models.Classstats

	err = database.GetDBEngine().Where("accId = ?", account.Id).Find(&classes)

	if utils.DefaultErrorHandler(ctx, err, logger){
		return
	}

	dailyQuest := database.GetDBEngine().Where("accId = ?").Find()

	xmlt := AccountXML{
		Id: account.Id,
		Name: account.Name,
		Namechosen: account.Namechosen,
		Admin: admin,
		Verified: verifiedEmail,
		Credits: stats.Credits,
		FortuneTokens: stats.Fortunetokens,
		NextCharSlotPrice: nextCharSlotPrice(account),
		BeginnerPackageTimeLeft: 1,
		PetYardType: account.Petyardtype,
		ArenaTickets: 0,
		Stats: statsXML{
			ClassStats: *toClassStatsXML(&classes),
		},
		DailyQuest: dailyQuestXML{

		}
	}

	ctx.XML(xmlt)
}

func toClassStatsXML(classstats *[]models.Classstats) *[]classStatsXML {
	var class []classStatsXML

	for _, v := range *classstats {
		class = append(class, classStatsXML{
			ObjectType: fmt.Sprintf("%d",v.Objtype),
			BestLevel: v.Bestlv,
			BestFame: v.Bestfame,
		})
	}

	return &class
}
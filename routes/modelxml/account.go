package modelxml

import (
	"encoding/xml"
	"fmt"
	"hagnix-server-go1/database/models"
)

type VaultXML struct {
	XMLName xml.Name `xml:"Vault"`
	Chests  []string `xml:"Chest"`
}

type GuildXML struct {
	XMLName xml.Name `xml:"Guild"`
	Id      int      `xml:"id,attr"`
	Rank    int      `xml:"Rank"`
	Name    string   `xml:"Name"`
	Fame    int      `xml:"Fame"`
}

type GiftsXML struct {
	XMLName xml.Name `xml:"Gifts"`
	Gifts   string   `xml:",innerxml"`
}

type DailyQuestXML struct {
	XMLName     xml.Name `xml:"DailyQuest"`
	Description string   `xml:"Description"`
	Image       string   `xml:"Image"`
	Tier        int      `xml:"tier,attr"`
	Goal        int      `xml:"goal,attr"`
}

type StatsXML struct {
	XMLName      xml.Name `xml:"Stats"`
	ClassStats   []ClassStatsXML
	BestCharFame int `xml:"BestCharFame"`
	TotalFame    int `xml:"TotalFame"`
	Fame         int `xml:"Fame"`
}

type ClassStatsXML struct {
	XMLName    xml.Name `xml:"ClassStats"`
	ObjectType string   `xml:"objectType,attr"`
	BestLevel  int      `xml:"BestLevel"`
	BestFame   int      `xml:"BestFame"`
}

type AccountXML struct {
	XMLName                 xml.Name `xml:"Account"`
	Id                      int64    `xml:"AccountId"`
	Name                    string   `xml:"Name"`
	Namechosen              bool     `xml:"NameChosen"`
	Admin                   bool     `xml:"Admin"`
	Verified                bool     `xml:"VerifiedEmail"`
	Credits                 int      `xml:"Credits"`
	FortuneTokens           int      `xml:"FortuneTokens"`
	NextCharSlotPrice       int      `xml:"NextCharSlotPrice"`
	BeginnerPackageTimeLeft int      `xml:"BeginnerPackageTimeLeft"`
	PetYardType             int      `xml:"PetYardType"`
	ArenaTickets            int      `xml:"ArenaTickets"`
	IsAgeVerified           int      `xml:"IsAgeVerified"`
	Banned                  bool     `xml:"Banned"`
	Stats                   StatsXML
	DailyQuest              DailyQuestXML
	Guild                   GuildXML
	Gifts                   GiftsXML
	Vault                   VaultXML
}

var DailyQuestDescriptions = []string{
	"This is the first quest of the day! Bring me a {goal} and I will reward you with a fortune token! But if you can complete all the quests, there will be an added bonus for you!",
	"Ahh, you have moved on to the second quest! If you bring me a {goal} I can pull out the magic bits and make another Fortune Token! If you finish my next quest, I will up the ante a bit...",
	"You again! Excellent Since you have been so helpful, I will use some specific parts arround here and make you TWO Fortune Tokens. All I need is a {goal}",
	"You again! Excellent Since you have been so helpful, I will use some specific parts arround here and make you TWO Fortune Tokens. All I need is a {goal}",
}

var ImageTiers = []string{
	"http://rotmg.kabamcdn.com/DailyQuest1FortuneToken.png",
	"http://rotmg.kabamcdn.com/DailyQuest1FortuneToken.png",
	"http://rotmg.kabamcdn.com/DailyQuest2FortuneToken.png",
	"http://rotmg.kabamcdn.com/DailyQuest2FortuneToken.png",
}

func ToVaultXML(vaults []models.Vaults) *VaultXML {
	vault := VaultXML{}

	for _, v := range vaults {
		vault.Chests = append(vault.Chests, v.Items)
	}

	return &vault
}

func ToClassStatsXML(classstats []models.Classstats) []ClassStatsXML {
	var class []ClassStatsXML

	for _, v := range classstats {
		class = append(class, ClassStatsXML{
			ObjectType: fmt.Sprintf("0x%x", v.Objtype),
			BestLevel:  v.Bestlv,
			BestFame:   v.Bestfame,
		})
	}

	return class
}

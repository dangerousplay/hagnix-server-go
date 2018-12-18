package modelxml

import (
	"encoding/xml"
	"github.com/ivahaev/go-logger"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

var cache *ItemsXML = &ItemsXML{}

type ObjectXML struct {
	Text    string `xml:",chardata"`
	Type    string `xml:"type,attr"`
	ID      string `xml:"id,attr"`
	Class   string `xml:"Class"`
	Item    string `xml:"Item"`
	Texture struct {
		Text  string `xml:",chardata"`
		File  string `xml:"File"`
		Index string `xml:"Index"`
	} `xml:"Texture"`
	SlotType        string `xml:"SlotType"`
	Description     string `xml:"Description"`
	ActivateOnEquip []struct {
		Text   string `xml:",chardata"`
		Stat   string `xml:"stat,attr"`
		Amount string `xml:"amount,attr"`
	} `xml:"ActivateOnEquip"`
	BagType    string `xml:"BagType"`
	FameBonus  string `xml:"FameBonus"`
	FeedPower  string `xml:"feedPower"`
	Soulbound  string `xml:"Soulbound"`
	DisplayId  string `xml:"DisplayId"`
	UnlockCost int    `xml:"UnlockCost"`
}

type ItemsXML struct {
	XMLName xml.Name    `xml:"Objects"`
	Text    string      `xml:",chardata"`
	Object  []ObjectXML `xml:"Object"`
}

func GetItems() *ItemsXML {
	return cache
}

func GetItem(typ string) *ObjectXML {
	var item *ObjectXML
	for _, v := range cache.Object {
		if v.Type == typ {
			item = &v
			break
		}
	}
	return item
}

func GetBonusItem(typ string) int {
	value, _ := strconv.Atoi(GetItem(typ).FameBonus)

	return value
}

func InitItems() {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	file, err := os.Open(path + "/game/dat1.xml")

	if err != nil {
		logger.Warnf("Can't open dat1.xml: %s", err.Error())
		return
	}

	readed, err := ioutil.ReadAll(file)

	if err != nil {
		logger.Warnf("Can't read file dat1.xml: %s", err.Error())
		return
	}

	err = xml.Unmarshal(readed, cache)

	if err != nil || cache == nil || len(cache.Object) < 1 {
		logger.Warnf("Erro Marshaling dat1.xml: %s", err.Error())
		return
	}
}

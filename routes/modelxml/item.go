package modelxml

import "encoding/xml"

var Items = []ItemCostXML{
	{Type: "900", Purchasable: 0, Expires: 0, Price: "90000"},
	{Type: "902", Purchasable: 0, Expires: 0, Price: "90000"},
	{Type: "834", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "835", Purchasable: 1, Expires: 0, Price: "600"},
	{Type: "836", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "837", Purchasable: 1, Expires: 0, Price: "600"},
	{Type: "838", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "839", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "840", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "841", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "842", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "843", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "844", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "845", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "846", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "847", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "848", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "849", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "850", Purchasable: 0, Expires: 1, Price: "900"},
	{Type: "851", Purchasable: 0, Expires: 1, Price: "900"},
	{Type: "852", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "853", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "854", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "855", Purchasable: 1, Expires: 0, Price: "900"},
	{Type: "856", Purchasable: 0, Expires: 0, Price: "90000"},
	{Type: "883", Purchasable: 0, Expires: 0, Price: "90000"},
}

var MaxClassLevels = []MaxClassLevelItem{
	{ClassType: "768", MaxLevel: "20"},
	{ClassType: "800", MaxLevel: "20"},
	{ClassType: "802", MaxLevel: "20"},
	{ClassType: "803", MaxLevel: "20"},
	{ClassType: "804", MaxLevel: "20"},
	{ClassType: "805", MaxLevel: "20"},
	{ClassType: "806", MaxLevel: "20"},
	{ClassType: "775", MaxLevel: "20"},
	{ClassType: "782", MaxLevel: "20"},
	{ClassType: "797", MaxLevel: "20"},
	{ClassType: "784", MaxLevel: "20"},
	{ClassType: "801", MaxLevel: "20"},
	{ClassType: "798", MaxLevel: "20"},
	{ClassType: "799", MaxLevel: "20"},
}

type MaxClassLevelItem struct {
	XMLName   xml.Name `xml:"MaxClassLevel"`
	ClassType string   `xml:"classType,attr"`
	MaxLevel  string   `xml:"maxLevel,attr"`
}

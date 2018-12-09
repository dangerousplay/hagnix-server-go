package modelxml

import "encoding/xml"

type BonusXML struct {
	XMLName     xml.Name `xml:"Bonus"`
	Id          string   `xml:"id,attr"`
	Description string   `xml:"desc,attr"`
	Value       string   `xml:",innerxml"`
}

type DeathXML struct {
	XMLName   xml.Name `xml:"Fame"`
	Name      string   `xml:"Account>Name"`
	BaseFame  int      `xml:"BaseFame"`
	Bonus     []BonusXML
	TotalFame int    `xml:"TotalFame"`
	CreatedOn string `xml:"CreatedOn"`
	KilledBy  string `xml:"KilledBy"`
}

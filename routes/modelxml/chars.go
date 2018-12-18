package modelxml

import "encoding/xml"

var Classes = []ClassAvailabilityXML{
	{Class: "Rogue", Restricted: "restricted"},
	{Class: "Assassin", Restricted: "restricted"},
	{Class: "Huntress", Restricted: "restricted"},
	{Class: "Mystic", Restricted: "restricted"},
	{Class: "Trickster", Restricted: "restricted"},
	{Class: "Sorcerer", Restricted: "restricted"},
	{Class: "Ninja", Restricted: "restricted"},
	{Class: "Archer", Restricted: "restricted"},
	{Class: "Wizard", Restricted: "unrestricted"},
	{Class: "Priest", Restricted: "restricted"},
	{Class: "Necromancer", Restricted: "restricted"},
	{Class: "Warrior", Restricted: "restricted"},
	{Class: "Knight", Restricted: "restricted"},
	{Class: "Paladin", Restricted: "restricted"},
}

type AbilityItemXML struct {
	XMLName xml.Name `xml:"Ability"`
	Type    int      `xml:"type,attr"`
	Power   int      `xml:"power,attr"`
	Points  int      `xml:"points,attr"`
}

type AbilityWrapper struct{ Abilities []AbilityItemXML }

type PetItemXML struct {
	SkinName        string         `xml:"name,attr"`
	Type            int            `xml:"type,attr"`
	InstanceId      int            `xml:"instanceId,attr"`
	MaxAbilityPower int            `xml:"maxAbilityPower,attr"`
	Skin            int            `xml:"skin,attr"`
	Rarity          int            `xml:"rarity,attr"`
	Abilities       AbilityWrapper `xml:"Abilities"`
}

type ItemXML struct {
	XMLName xml.Name `xml:"Item"`
	Icon    string   `xml:"Icon"`
	Title   string   `xml:"Title"`
	TagLine string   `xml:"TagLine"`
	Link    string   `xml:"Link"`
	Date    int      `xml:"Date"`
}

type NewsItemXML struct {
	XMLName xml.Name `xml:"Item"`
	Icon    string   `xml:"Icon"`
	Title   string   `xml:"Title"`
	TagLine string   `xml:"TagLine"`
	Link    string   `xml:"Link"`
	Date    int64    `xml:"Date"`
}

type ServerItemXML struct {
	XMLName      xml.Name `xml:"Server"`
	Name         string   `xml:"Name"`
	DNS          string   `xml:"DNS"`
	Lat          float32  `xml:"Lat"`
	Long         float32  `xml:"Long"`
	Usage        float64  `xml:"Usage"`
	RankRequired int      `xml:"RankRequired"`
	AdminOnly    string   `xml:"AdminOnly"`
}

type CharXML struct {
	XMLName          xml.Name    `xml:"Char"`
	Name             string      `xml:"Account>Name,omitempty"`
	Id               int         `xml:"id,attr"`
	ObjectType       int         `xml:"ObjectType"`
	Level            int         `xml:"Level"`
	Exp              int         `xml:"Exp"`
	CurrentFame      int         `xml:"CurrentFame"`
	HealthStackCount int         `xml:"HealthStackCount"`
	MagicStackCount  int         `xml:"MagicStackCount"`
	Equipment        string      `xml:"Equipment"`
	HasBackpack      int         `xml:"HasBackpack"`
	MaxHitPoints     int         `xml:"MaxHitPoints"`
	HitPoints        int         `xml:"HitPoints"`
	MaxMagicPoints   int         `xml:"MaxMagicPoints"`
	MagicPoints      int         `xml:"MagicPoints"`
	Attack           int         `xml:"Attack"`
	Defense          int         `xml:"Defense"`
	Speed            int         `xml:"Speed"`
	Dexterity        int         `xml:"Dexterity"`
	HpRegen          int         `xml:"HpRegen"`
	MpRegen          int         `xml:"MpRegen"`
	Tex1             int         `xml:"Tex1"`
	Tex2             int         `xml:"Tex2"`
	XpBoosted        bool        `xml:"XpBoosted"`
	XpTimer          int         `xml:"XpTimer"`
	LDTimer          int         `xml:"LDTimer"`
	LTTimer          int         `xml:"LTTimer"`
	PCStats          string      `xml:"PCStats"`
	CasToken         string      `xml:"casToken,omitempty"`
	Skin             int         `xml:"Texture"`
	Dead             bool        `xml:"Dead"`
	Pet              *PetItemXML `xml:"Pet,omitempty"`
}

type ClassAvailabilityXML struct {
	XMLName    xml.Name `xml:"ClassAvailability"`
	Class      string   `xml:"id,attr"`
	Restricted string   `xml:",innerxml"`
}

type ItemCostXML struct {
	Type        string `xml:"type,attr"`
	Purchasable int    `xml:"purchasable,attr"`
	Expires     int    `xml:"expires,attr"`
	Price       string `xml:",innerxml"`
}

type CharsXML struct {
	XMLName           xml.Name               `xml:"Chars"`
	Char              []CharXML              `xml:"Char"`
	NextCharId        int                    `xml:"nextCharId,attr"`
	MaxNumChars       int                    `xml:"maxNumChars,attr"`
	Account           AccountXML             `xml:"Account"`
	NewsXML           []NewsItemXML          `xml:"News>Item"`
	Servers           []ServerItemXML        `xml:"Servers>Server"`
	OwnedSkins        string                 `xml:"OwnedSkins,omitempty"`
	TOSPopup          *int                   `xml:"TOSPopup"`
	Lat               string                 `xml:"Lat,omitempty"`
	Long              string                 `xml:"Long,omitempty"`
	Classes           []ClassAvailabilityXML `xml:"ClassAvailabilityList>ClassAvailability"`
	ItemCosts         []ItemCostXML          `xml:"ItemCosts>ItemCost"`
	MaxClassLevelList []MaxClassLevelItem    `xml:"MaxClassLevelList>MaxClassLevel"`
}

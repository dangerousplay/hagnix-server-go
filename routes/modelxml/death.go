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
	Name      string   `xml:"Account>Name,omitempty"`
	BaseFame  int      `xml:"BaseFame"`
	Character CharXML
	Bonus     []BonusXML
	TotalFame int    `xml:"TotalFame"`
	CreatedOn string `xml:"CreatedOn"`
	KilledBy  string `xml:"KilledBy"`

	Shots                  int32
	ShotsThatDamage        int32
	SpecialAbilityUses     int32
	TilesUncovered         int32
	Teleports              int32
	PotionsDrunk           int32
	MonsterKills           int32
	MonsterAssists         int32
	GodKills               int32
	GodAssists             int32
	CubeKills              int32
	OryxKills              int32
	QuestsCompleted        int32
	PirateCavesCompleted   int32
	UndeadLairsCompleted   int32
	AbyssOfDemonsCompleted int32
	SnakePitsCompleted     int32
	SpiderDensCompleted    int32
	SpriteWorldsCompleted  int32
	LevelUpAssists         int32
	MinutesActive          int32
	TombsCompleted         int32
	TrenchesCompleted      int32
	JunglesCompleted       int32
	ManorsCompleted        int32
}

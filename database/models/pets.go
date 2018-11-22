package models

type Pets struct {
	Accid     int    `xorm:"not null pk INT(11)"`
	Petid     int    `xorm:"not null pk autoincr INT(11)"`
	Objtype   int    `xorm:"not null SMALLINT(5)"`
	Skinname  string `xorm:"not null VARCHAR(128)"`
	Skin      int    `xorm:"not null INT(11)"`
	Family    int    `xorm:"not null INT(11)"`
	Rarity    int    `xorm:"not null INT(11)"`
	Maxlevel  int    `xorm:"not null default 30 INT(11)"`
	Abilities string `xorm:"not null VARCHAR(128)"`
	Levels    string `xorm:"not null VARCHAR(128)"`
	Xp        string `xorm:"not null default '0, 0, 0' VARCHAR(128)"`
}

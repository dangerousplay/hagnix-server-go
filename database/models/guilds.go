package models

type Guilds struct {
	Id             int    `xorm:"not null pk autoincr INT(11)"`
	Name           string `xorm:"not null default 'DEFAULT_GUILD' VARCHAR(128)"`
	Members        string `xorm:"not null pk VARCHAR(128)"`
	Guildfame      int    `xorm:"not null INT(11)"`
	Totalguildfame int    `xorm:"not null INT(11)"`
	Level          int    `xorm:"not null default 1 INT(11)"`
}

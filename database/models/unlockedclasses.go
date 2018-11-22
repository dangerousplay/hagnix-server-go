package models

type Unlockedclasses struct {
	Id        int    `xorm:"not null pk autoincr INT(11)"`
	Accid     int    `xorm:"not null INT(11)"`
	Class     string `xorm:"not null VARCHAR(128)"`
	Available string `xorm:"not null VARCHAR(128)"`
}

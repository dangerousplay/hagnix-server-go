package models

type Backpacks struct {
	Accid  int    `xorm:"not null pk INT(11)"`
	Charid int    `xorm:"not null pk INT(11)"`
	Items  string `xorm:"not null default '-1, -1, -1, -1, -1, -1, -1, -1' VARCHAR(128)"`
}

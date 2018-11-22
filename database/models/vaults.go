package models

type Vaults struct {
	Accid   int    `xorm:"not null pk INT(11)"`
	Chestid int    `xorm:"not null pk autoincr INT(11)"`
	Items   string `xorm:"not null VARCHAR(128)"`
}

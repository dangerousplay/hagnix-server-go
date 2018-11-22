package models

type Giftcodes struct {
	Code    string `xorm:"not null pk VARCHAR(128)"`
	Content string `xorm:"not null VARCHAR(512)"`
	Accid   int    `xorm:"not null default 0 INT(11)"`
}

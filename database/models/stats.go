package models

type Stats struct {
	Accid              int `xorm:"not null pk INT(11)"`
	Fame               int `xorm:"not null pk default 0 INT(11)"`
	Totalfame          int `xorm:"not null pk default 0 INT(11)"`
	Credits            int `xorm:"not null default 0 INT(11)"`
	Totalcredits       int `xorm:"not null default 0 INT(11)"`
	Fortunetokens      int `xorm:"not null default 0 INT(11)"`
	Totalfortunetokens int `xorm:"not null default 0 INT(11)"`
}

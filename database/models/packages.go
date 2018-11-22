package models

import (
	"time"
)

type Packages struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	Name        string    `xorm:"not null VARCHAR(128)"`
	Maxpurchase int       `xorm:"not null default -1 INT(11)"`
	Weight      int       `xorm:"not null default 0 INT(11)"`
	Contents    string    `xorm:"not null TEXT"`
	Bgurl       string    `xorm:"not null VARCHAR(512)"`
	Price       int       `xorm:"not null INT(11)"`
	Quantity    int       `xorm:"not null default -1 INT(11)"`
	Enddate     time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
}

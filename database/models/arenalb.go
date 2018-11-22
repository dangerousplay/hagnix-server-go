package models

import (
	"time"
)

type Arenalb struct {
	Id     int       `xorm:"not null pk autoincr INT(11)"`
	Wave   int       `xorm:"not null pk INT(11)"`
	Accid  int       `xorm:"not null INT(11)"`
	Charid int       `xorm:"not null INT(11)"`
	Petid  int       `xorm:"INT(11)"`
	Time   string    `xorm:"not null VARCHAR(256)"`
	Date   time.Time `xorm:"not null DATETIME"`
}

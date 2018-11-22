package models

import (
	"time"
)

type Globalnews struct {
	Slot       int       `xorm:"not null pk INT(11)"`
	Linktype   int       `xorm:"not null INT(11)"`
	Title      string    `xorm:"not null VARCHAR(65)"`
	Image      string    `xorm:"not null TEXT"`
	Priority   int       `xorm:"not null INT(11)"`
	Linkdetail string    `xorm:"not null TEXT"`
	Platform   string    `xorm:"not null default 'kabam.com,kongregate,steam,rotmg' VARCHAR(128)"`
	Starttime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Endtime    time.Time `xorm:"not null default '0000-00-00 00:00:00' TIMESTAMP"`
}

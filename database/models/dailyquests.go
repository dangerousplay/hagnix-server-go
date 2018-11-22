package models

import (
	"time"
)

type Dailyquests struct {
	Accid int       `xorm:"not null pk INT(11)"`
	Goals string    `xorm:"not null VARCHAR(512)"`
	Tier  int       `xorm:"not null default 1 INT(11)"`
	Time  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

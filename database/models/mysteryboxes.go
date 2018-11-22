package models

import (
	"time"
)

type Mysteryboxes struct {
	Id            int       `xorm:"not null pk INT(11)"`
	Title         string    `xorm:"not null VARCHAR(128)"`
	Weight        int       `xorm:"not null INT(11)"`
	Description   string    `xorm:"not null VARCHAR(128)"`
	Contents      string    `xorm:"not null TEXT"`
	Priceamount   int       `xorm:"not null INT(11)"`
	Pricecurrency int       `xorm:"not null INT(11)"`
	Image         string    `xorm:"not null TEXT"`
	Icon          string    `xorm:"not null TEXT"`
	Saleprice     int       `xorm:"not null INT(11)"`
	Salecurrency  int       `xorm:"not null INT(11)"`
	Saleend       time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	Starttime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Boxend        time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
}

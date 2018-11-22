package models

import (
	"time"
)

type Thealchemist struct {
	Id                int       `xorm:"not null pk INT(11)"`
	Title             string    `xorm:"not null VARCHAR(512)"`
	Description       string    `xorm:"VARCHAR(512)"`
	Image             string    `xorm:"not null VARCHAR(512)"`
	Icon              string    `xorm:"not null VARCHAR(512)"`
	Contents          string    `xorm:"not null TEXT"`
	Pricefirstingold  int       `xorm:"not null default 51 INT(11)"`
	Pricefirstintoken int       `xorm:"not null default 1 INT(11)"`
	Pricesecondingold int       `xorm:"not null default 75 INT(11)"`
	Starttime         time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Endtime           time.Time `xorm:"not null DATETIME"`
}

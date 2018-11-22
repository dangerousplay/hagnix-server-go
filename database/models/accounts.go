package models

import (
	"time"
)

type Accounts struct {
	Id             int64     `xorm:"pk autoincr BIGINT(255)"`
	Uuid           string    `xorm:"not null pk VARCHAR(128)"`
	ObjectId       string    `xorm:"not null VARCHAR(30)"`
	Password       string    `xorm:"not null VARCHAR(256)"`
	Name           string    `xorm:"not null default 'DEFAULT' VARCHAR(64)"`
	Rank           int       `xorm:"not null default 0 INT(1)"`
	Namechosen     int       `xorm:"not null default 0 TINYINT(1)"`
	Verified       int       `xorm:"not null default 1 TINYINT(1)"`
	Guild          int       `xorm:"not null pk INT(11)"`
	Guildrank      int       `xorm:"not null INT(11)"`
	Guildfame      int       `xorm:"not null default 0 INT(11)"`
	Lastip         string    `xorm:"not null pk default '' VARCHAR(128)"`
	Vaultcount     int       `xorm:"not null default 1 INT(11)"`
	Maxcharslot    int       `xorm:"not null default 2 INT(11)"`
	Regtime        time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	Guest          int       `xorm:"not null default 0 TINYINT(1)"`
	Banned         int       `xorm:"not null pk default 0 TINYINT(1)"`
	Publicmuledump int       `xorm:"not null default 1 INT(1)"`
	Muted          int       `xorm:"not null default 0 TINYINT(1)"`
	Prodacc        int       `xorm:"not null default 0 TINYINT(1)"`
	Locked         string    `xorm:"not null VARCHAR(512)"`
	Ignored        string    `xorm:"not null VARCHAR(512)"`
	Gifts          string    `xorm:"not null default '' VARCHAR(10000)"`
	Isageverified  int       `xorm:"not null default 0 TINYINT(1)"`
	Petyardtype    int       `xorm:"not null default 1 INT(11)"`
	Ownedskins     string    `xorm:"not null default '' VARCHAR(2048)"`
	Authtoken      string    `xorm:"not null default '' VARCHAR(128)"`
	Acceptednewtos int       `xorm:"not null default 1 TINYINT(1)"`
	Lastseen       time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	Accountinuse   int       `xorm:"not null default 0 TINYINT(1)"`
}

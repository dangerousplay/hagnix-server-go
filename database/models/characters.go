package models

import (
	"time"
)

type Characters struct {
	Id            int       `xorm:"not null pk autoincr INT(11)"`
	Accid         int       `xorm:"not null pk INT(11)"`
	Charid        int       `xorm:"not null INT(11)"`
	Chartype      int       `xorm:"not null default 782 INT(11)"`
	Level         int       `xorm:"not null default 1 INT(11)"`
	Exp           int       `xorm:"not null default 0 INT(11)"`
	Fame          int       `xorm:"not null default 0 INT(11)"`
	Items         string    `xorm:"not null default '-1, -1, -1, -1' VARCHAR(128)"`
	Hppotions     int       `xorm:"not null default 0 INT(11)"`
	Mppotions     int       `xorm:"not null default 0 INT(11)"`
	Hp            int       `xorm:"not null default 1 INT(11)"`
	Mp            int       `xorm:"not null default 1 INT(11)"`
	Stats         string    `xorm:"not null default '1, 1, 1, 1, 1, 1, 1, 1' VARCHAR(128)"`
	Dead          int       `xorm:"not null pk default 0 TINYINT(1)"`
	Tex1          int       `xorm:"not null default 0 INT(11)"`
	Tex2          int       `xorm:"not null default 0 INT(11)"`
	Pet           int       `xorm:"not null default -1 INT(11)"`
	Petid         int       `xorm:"not null default -1 INT(11)"`
	Hasbackpack   int       `xorm:"not null default 0 INT(11)"`
	Skin          int       `xorm:"not null default 0 INT(11)"`
	Xpboostertime int       `xorm:"not null default 0 INT(11)"`
	Ldtimer       int       `xorm:"not null default 0 INT(11)"`
	Lttimer       int       `xorm:"not null default 0 INT(11)"`
	Famestats     string    `xorm:"not null VARCHAR(512)"`
	Createtime    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Deathtime     time.Time `xorm:"not null default '0000-00-00 00:00:00' TIMESTAMP"`
	Totalfame     int       `xorm:"not null default 0 INT(11)"`
	Lastseen      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Lastlocation  string    `xorm:"not null default '' VARCHAR(128)"`
}

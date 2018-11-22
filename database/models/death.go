package models

import (
	"time"
)

type Death struct {
	AccountId   int       `xorm:"not null pk unique INT(11) 'AccId'"`
	CharacterId int       `xorm:"not null pk unique INT(11) 'chrId'"`
	Name        string    `xorm:"not null default 'DEFAULT' unique VARCHAR(64)"`
	CharType    int       `xorm:"not null default 782 unique INT(11) 'charType'"`
	Tex1        int       `xorm:"not null default 0 INT(11)"`
	Tex2        int       `xorm:"not null default 0 INT(11)"`
	Skin        int       `xorm:"not null default 0 INT(11)"`
	Items       string    `xorm:"not null default '-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1' VARCHAR(128)"`
	Fame        int       `xorm:"not null default 0 INT(11)"`
	Exp         int       `xorm:"not null INT(11)"`
	Famestats   string    `xorm:"not null VARCHAR(256)"`
	TotalFame   int       `xorm:"not null default 0 INT(11) 'totalFame'"`
	Firstborn   int       `xorm:"not null TINYINT(1)"`
	Killer      string    `xorm:"not null VARCHAR(128)"`
	Time        time.Time `xorm:"not null pk default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

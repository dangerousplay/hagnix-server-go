package models

import (
	"time"
)

type News struct {
	Id    int       `xorm:"not null pk autoincr INT(11)"`
	Icon  string    `xorm:"not null default 'info' VARCHAR(16)"`
	Title string    `xorm:"not null default 'Default news title' VARCHAR(128)"`
	Text  string    `xorm:"not null pk default 'Default news text' VARCHAR(128)"`
	Link  string    `xorm:"not null default 'http://mmoe.net/' VARCHAR(256)"`
	Date  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

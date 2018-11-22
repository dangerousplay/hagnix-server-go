package models

type Classstats struct {
	Accid    int `xorm:"not null pk INT(11)"`
	Objtype  int `xorm:"not null pk INT(11)"`
	Bestlv   int `xorm:"not null default 1 INT(11)"`
	Bestfame int `xorm:"not null default 0 INT(11)"`
}

package models

type Boards struct {
	Guildid int    `xorm:"not null pk INT(11)"`
	Text    string `xorm:"not null VARCHAR(1024)"`
}

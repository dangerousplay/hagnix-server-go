package database

type Death struct {
	AccountId   int    `xorm:"int(11) not null unique 'accId'"`
	CharacterId int    `xorm:"int(11) not null unique 'chrId'"`
	Name        string `xorm:"varchar(64) unique 'name'"`
	CharType    int    `xorm:"int(11) not null unique 'charType'"`
	Text1       int    `xorm:"int(11) 'text1'"`
	Text2       int    `xorm:"int(11) 'text2'"`
	Skin        int    `xorm:"int(11) 'skin'"`
	Items       string `xorm:"varchar(128) 'items'"`
	Fame        int    `xorm:"int(11) 'fame'"`
	Exp         int    `xorm:"int(11) 'exp'"`
	FameStats   string `xorm:"varchar(256) 'fameStats'"`
	TotalFame   int    `xorm:"int(11) 'totalFame'"`
	FirstBorn   int    `xorm:"tinyint(1) 'firstBorn'"`
	Killer      string `xorm:"varchar(128) 'killer'"`
	TimeStramp  string `xorm:"TIMESTAMP 'time'"`
}

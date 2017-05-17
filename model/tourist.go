package model

type Tourist struct {
	Id           int    `xorm:"not null pk autoincr INT(11)"`
	Title        string `xorm:"default '' VARCHAR(500)"`
	Introduction string `xorm:"default '' VARCHAR(1000)"`
	Img          string `xorm:"not null VARCHAR(500)"`
	Detail       string `xorm:"VARCHAR(100)"`
	Collected    int    `xorm:"not null default 0 TINYINT(1)"`
}

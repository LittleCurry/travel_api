package model

type LifeDetail struct {
	Id     int    `xorm:"not null pk autoincr unique INT(11)"`
	LifeId int    `xorm:"not null INT(11)"`
	Img    []byte `xorm:"LONGBLOB"`
}

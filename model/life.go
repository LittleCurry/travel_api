package model

type Life struct {
	Id   int    `xorm:"not null pk autoincr unique INT(11)"`
	Info string `xorm:"VARCHAR(200)"`
	Img  []byte `xorm:"not null LONGBLOB"`
}

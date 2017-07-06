package model

import (
	"time"
)

type Life struct {
	Id         int       `xorm:"not null pk autoincr unique INT(11)"`
	Info       string    `xorm:"VARCHAR(200)"`
	CreateTime time.Time `xorm:"DATETIME"`
}

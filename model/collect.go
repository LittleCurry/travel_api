package model

import (
	"time"
)

type Collect struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Img        string    `xorm:"VARCHAR(200)"`
	Title      string    `xorm:"default '' VARCHAR(50)"`
	Money      int       `xorm:"default 0 INT(11)"`
	CreateDate time.Time `xorm:"DATETIME"`
	Src        string    `xorm:"default '' VARCHAR(100)"`
	TouristId  int       `xorm:"not null INT(11)"`
}

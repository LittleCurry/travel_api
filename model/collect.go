package model

type Collect struct {
	Id    int    `xorm:"not null pk autoincr INT(11)"`
	Title string `xorm:"default '' VARCHAR(50)"`
	Img   string `xorm:"VARCHAR(200)"`
	Src   string `xorm:"default '' VARCHAR(100)"`
}

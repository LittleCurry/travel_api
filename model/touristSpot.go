package model

type TouristSpot struct {
	Id           int    `xorm:"not null pk autoincr INT(11)"`
	Title        string `xorm:"default '' VARCHAR(200)"`
	Introduction string `xorm:"default '' VARCHAR(200)"`
	Img          string `xorm:"not null VARCHAR(500)"`
}

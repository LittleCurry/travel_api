package joinmodel

import "git.iguiyu.com/park/struct/model"

type LockerCharging struct {
	model.Locker `xorm:"extends"`
	MonthlyFee   int `xorm:"INT(11)"`
	Commission   int `xorm:"TINYINT(4)"`
}

func (LockerCharging) TableName() string {
	return "locker"
}

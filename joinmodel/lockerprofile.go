package joinmodel

import "git.iguiyu.com/park/struct/model"

type LockerProfile struct {
	model.Locker `xorm:"extends"`
	Pname        string `xorm:"VARCHAR(40)"`
}

func (LockerProfile) TableName() string {
	return "locker"
}

package joinmodel

import "git.iguiyu.com/park/struct/model"

type ProfileDetailProfile struct {
	model.OwnerProfileDetail `xorm:"extends"`
	OwnerUserId              int `xorm:"not null index INT(11)"`
}

func (ProfileDetailProfile) TableName() string {
	return "owner_profile_detail"
}

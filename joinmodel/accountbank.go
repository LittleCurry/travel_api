package joinmodel

import "git.iguiyu.com/park/struct/model"

type AccountBank struct {
	model.Account `xorm:"extends"`
	BankName      string `xorm:"VARCHAR(45)"`
	BankCode      string `xorm:"VARCHAR(10)"`
	BankColor     string `xorm:"VARCHAR(10)"`
}

func (AccountBank) TableName() string {
	return "account"
}

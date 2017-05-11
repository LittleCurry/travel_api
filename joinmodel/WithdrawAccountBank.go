package joinmodel

import "git.iguiyu.com/park/struct/model"

type WithdrawAccountBank struct {
	model.Withdraw `xorm:"extends"`
	AccountCode    string `xorm:"not null VARCHAR(100)"`
	AccountName    string `xorm:"not null VARCHAR(100)"`
	BankName       string `xorm:"VARCHAR(45)"`
	BankCode       string `xorm:"VARCHAR(10)"`
}

func (WithdrawAccountBank) TableName() string {
	return "withdraw"
}

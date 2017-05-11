package vm

import (
	"time"
)

type TxBankType struct {
	Id        int    `json:"bankid"`
	BankName  string `json:"bankname"`
	BankCode  string `json:"bankcode"`
	BankColor string `json:"bankcolor"`
}

type RxBankAccount struct {
	AccountName string `json:"accountname"`
	AccountCode string `json:"accountcode"`
	BankId      int    `json:"bankid"`
}

type TxBankAccount struct {
	Id          int    `json:"id"`
	AccountName string `json:"accountname"`
	AccountCode string `json:"accountcode"`
	BankId      int    `json:"bankid"`
	BankName    string `json:"bankname"`
	BankCode    string `json:"bankcode"`
	BankColor   string `json:"bankcolor"`
}

type RxWithdraw struct {
	Money     int    `json:"money"`
	AccountId int    `json:"accountid"`
	Code      string `json:"code"`
}

type TxWithdraw struct {
	Id          int    `json:"id"`
	Money       int    `json:"money"`
	CreateDate  string `json:"createdate"`
	UpdateDate  string `json:"updatedate"`
	Status      int    `json:"status"`
	AccountCode string `json:"accountcode"`
	AccountName string `json:"accountname"`
	BankName    string `json:"bankname"`
	BankCode    string `json:"bankcode"`
}

func (withdraw *TxWithdraw) CreateTime(createTime time.Time) {
	withdraw.CreateDate = createTime.Format("2006-01-02")
}

func (withdraw *TxWithdraw) UpdateTime(updateTime time.Time) {
	withdraw.UpdateDate = updateTime.Format("2006-01-02")
}

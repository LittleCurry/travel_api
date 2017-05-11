package vm

import (
	"time"
)

type RxUser struct {
	LoginName string `json:"mobile"`
	Password  string `json:"password"`
}

type TxUser struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type TxMe struct {
	Nickname string `json:"name"`
	Type     int    `json:"type"`
	//Balance  int    `json:"balance"`
	CreateAt string `json:"createdAt"`
}

func (txMe *TxMe) CreateDate(createDate time.Time) {
	txMe.CreateAt = createDate.Format("2006-01-02")
}

type RxRegUser struct {
	Nickname  string `json:"name"`
	Loginname string `json:"mobile"`
	Passwd    string `json:"password"`
	Code      string `json:"code"`
}

type RxNewPassword struct {
	Loginname string `json:"mobile"`
	Passwd    string `json:"password"`
	Code      string `json:"code"`
}

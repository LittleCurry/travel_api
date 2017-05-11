package vm

import "time"

type TxMessage struct {
	Id      int    `json:"id"`
	Subject string `json:"subject"`
	Content string `json:"content"`
	Level   int    `json:"level"`
	Read    int    `json:"read"`
	Create  string `json:"createdate"`
}

func (this *TxMessage) CreateDate(createDate time.Time) {
	this.Create = createDate.Format("2006-01-02 15:04:05")
}

type TxUnRead struct {
	Count int `json:"count"`
}

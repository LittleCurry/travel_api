package vm

import "time"

type TxCollect struct {
	Id            	int	`json:"id"`
	Img 		string	`json:"img"`
	Title 		string	`json:"title"`
	Money  		int	`json:"money"`
	Create  	string  `json:"createdate"`
	Src 		string	`json:"src"`
}

func (this *TxCollect) CreateDate(createDate time.Time) {
	this.Create = createDate.Format("2006-01-02 15:04:05")
}

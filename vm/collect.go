package vm

import (
	"time"
	"fmt"
)

type TxCollect struct {
	Id            	int	`json:"id"`
	Img 		string	`json:"img"`
	Title 		string	`json:"title"`
	Money  		int	`json:"money"`
	Create  	string  `json:"createdate"`
	Src 		string	`json:"src"`
	TouristId       int	`json:"tourist_id"`
}

func (this *TxCollect) Setcreate() {
	this.Create = time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("时间是",this.Create)
}

//func (this *TxCollect) CreateDate() {
//	this.Create = time.Now().Format("2006-01-02 15:04:05")
//	fmt.Printf("时间是",this.Create)
//}

//func (this *TxCollect) CreateDate(createDate time.Time) {
//	this.Create = createDate.Format("2006-01-02 15:04:05")
//}

//func (this *TxCollect) TouristId(tourist_id int) {
//	this.Id = tourist_id
//}

//func (this *TxCollect) tourist_id(id int) {
//	this.TouristId = id
//	fmt.Printf("id是",this.TouristId)
//}

func (this *TxCollect) SettouristId(tourist_id int) {
	this.TouristId = tourist_id
}
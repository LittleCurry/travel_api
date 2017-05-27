package vm

type TxCollect struct {
	Id            	int	`json:"id"`
	Img 		string	`json:"img"`
	Title 		string	`json:"title"`
	Money  		int	`json:"money"`
	Create  	string  `json:"createdate"`
	Src 		string	`json:"src"`
	TouristId       int	`json:"tourist_id"`
}

/*
func (this *TxCollect) Detail(temp string) {
	this.Src = temp
	fmt.Printf("src是",this.Create)
}

func (this *TxCollect) Id(temp int){
	this.TouristId = temp
}

func NewTxCollect() *TxCollect {
	//fmt.Printf("时间是",Create)
	return &TxCollect{Create:time.Now().Format("2006-01-02 15:04:05")}
}
*/
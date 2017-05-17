package vm

type TxTourist struct {
	Id            	int	`json:"id"`
	Title 		string	`json:"title"`
	Introduction  	string	`json:"introduction"`
	Img 		string	`json:"img"`
	Detail 		string	`json:"detail"`
	Collected       int  	`json:"collected"`
}

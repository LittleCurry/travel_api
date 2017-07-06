package vm

import "time"

//type TxLife struct {
//	Id            	int			`json:"id"`
//	info 			string		`json:"info"`
//	Img  			[]byte		`json:"img"`
//}

type RxLife struct {
	Info  		string       `json:"info"`
	Imgs 		[]LifeDetail `json:"imgs"`
}

type LifeDetail struct {
	Img 	   []byte 	`json:"img"`
}

type TxLife struct {
	Id            	int				`json:"id"`
	Info 			string			`json:"info"`
	ShowTime 		string			`json:"create_time"`
	Imgs 			[]LifeDetail	`json:"imgs"`
}

func (this *TxLife) CreateTime(createDate time.Time) {
	this.ShowTime = createDate.Format("2006-01-02 15:04:05")
}
package vm

import "time"

type TxSummary struct {
	Income int `json:"income"`
	//Profit  int        `json:"profit"`
	//Fee     int        `json:"fee"`
	Balance int        `json:"balance"`
	Incomes []TxIncome `json:"reports"`
}

type TxIncome struct {
	Income    int    `json:"income"`
	CreatedAt string `json:"date"`
}

func (this *TxIncome) CreateDate(createDate time.Time) {
	this.CreatedAt = createDate.Format("2006-01-02")
}

type TxSummaryTotal struct {
	TodayIncome int `json:"todayincome"`
	Income      int `json:"income"`
	Profit      int `json:"profit"`
	Fee         int `json:"fee"`
	Balance     int `json:"balance"`
}

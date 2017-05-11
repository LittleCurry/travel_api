package vm

type RxLock struct {
	QRCode string `json:"qrcode"`
}

type TxUnActivatedLock struct {
	Id      int    `json:"id"`
	ShortId string `json:"code"`
}

type TxLock struct {
	Id             int    `json:"id"`
	OwnerProfileId int    `json:"planid"`
	ShortId        string `json:"code"`
	Pname          string `json:"planname"`
	Address        string `json:"address"`
	Longitude      string `json:"lon"`
	Latitude       string `json:"lat"`
}

type RxLockToBind struct {
	OwnerProfileId int    `json:"planid"`
	Address        string `json:"address"`
	Longitude      string `json:"lon"`
	Latitude       string `json:"lat"`
}

type TxLockForList struct {
	Id           int    `json:"id"`
	ShortId      string `json:"code"`
	Longitude    string `json:"lon"`
	Latitude     string `json:"lat"`
	Address      string `json:"address"`
	BatteryLevel int    `json:"batteryLevel"`
	Income       int    `json:"income"`
	Status       int    `json:"status"`
	Error        int    `json:"error"`
	//Available    int      `json:"available"`
}

type TxLockDetail struct {
	Id           int    `json:"id"`
	Status       int    `json:"status"`
	ShortId      string `json:"code"`
	Longitude    string `json:"lon"`
	Latitude     string `json:"lat"`
	Address      string `json:"address"`
	BatteryLevel int    `json:"batteryLevel"`
	Income       int    `json:"income"`
	Error        int    `json:"error"`
	Pname        string `json:"planname"`
	Available    int    `json:"available"`
}

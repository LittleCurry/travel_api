package vm

type TxPlan struct {
	Id            int          `json:"planId"`
	Pname         string       `json:"name"`
	Free          int          `json:"free"`
	ProfileDetail []PlanDetail `json:"timeIntervalPlans"`
}

type RxPlan struct {
	Pname      string       `json:"name"`
	Free       int          `json:"free"`
	PlanDetail []PlanDetail `json:"timeIntervalPlans"`
}

type PlanDetail struct {
	Fulltime     int `json:"isFullDay"`
	BeginTime    int `json:"startAt"`
	EndTime      int `json:"endAt"`
	PricePerHour int `json:"price"`
	CappedPrice  int `json:"cappedPrice"`
	Unit         int `json:"unit"`
}

/*
type RxPlanDetail struct {
	Fulltime     int    `json:"isFullDay"`
	BeginTime    int    `json:"startAt"`
	EndTime      int    `json:"endAt"`
	PricePerHour string `json:"price"`
	CappedPrice  string `json:"cappedPrice"`
	Unit         int    `json:"unit"`
}
*/

type RxPlanRename struct {
	Pname string `json:"name"`
}

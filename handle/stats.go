package handle

import (
	"net/http"
	"time"

	"git.iguiyu.com/park/api/db"
	"git.iguiyu.com/park/api/joinmodel"
	"git.iguiyu.com/park/api/vm"
	"git.iguiyu.com/park/struct/model"
	"github.com/labstack/echo"
	"strconv"
)

func SummaryStats(c echo.Context) (err error) {
	startAt := c.QueryParam("startAt")
	endAt := c.QueryParam("endAt")

	if startAt != "" && endAt != "" {
		userId := c.Get("userId").(string)
		//incomes := make([]model.Income, 0)
		user := model.OwnerUser{}
		has, err0 := db.MySQL().Id(userId).Get(&user)
		if err0 == nil && has {
			//err1 := db.MySQL().Where("owner_user_id=?", userId).And("create_date <= ? AND create_date >= ?", endAt, startAt).GroupBy() .Asc("create_date").Find(&incomes)
			sql := "SELECT sum(income) income, Date(create_date) create_date FROM `income` WHERE owner_user_id=" + userId + " AND create_date <= '" + endAt + "' AND create_date >= '" + startAt + "' group by Date(create_date) order by Date(create_date) ASC"
			results, err1 := db.MySQL().Query(sql)
			if err1 == nil {
				txSummary := vm.TxSummary{}
				for _, result := range results {
					i, _ := strconv.Atoi(string(result["income"]))
					txSummary.Income += i
					txSummary.Incomes = append(txSummary.Incomes, vm.TxIncome{Income: i, CreatedAt: string(result["create_date"])})
				}
				txSummary.Balance = user.Balance
				//copier.Copy(&txSummary.Incomes, &results)

				return c.JSON(http.StatusOK, txSummary)
			}
		}
	}

	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func BriefStats(c echo.Context) (err error) {
	userId := c.Get("userId").(string)
	user := model.OwnerUser{}
	has, err1 := db.MySQL().Id(userId).Get(&user)
	if err1 == nil && has {
		txSummaryTotal := vm.TxSummaryTotal{}
		lockerCharging := make([]joinmodel.LockerCharging, 0)
		err2 := db.MySQL().Join(
			"INNER",
			"charging_policy",
			"charging_policy.id = locker.charging_policy_id").Where(
			"locker.owner_user_id=?", userId).Find(&lockerCharging)
		if err2 == nil {
			total, err3 := db.MySQL().Where("owner_user_id=?", userId).And("DATE(create_date) = DATE(NOW())").Sum(&model.Income{}, "income")
			if err3 == nil {
				for _, v := range lockerCharging {
					txSummaryTotal.Income += v.Income
					txSummaryTotal.Fee += v.Income * v.Commission / 100
					month := time.Since(v.ActiveDate) / time.Hour / 365 / 2
					txSummaryTotal.Fee += int(month) * v.MonthlyFee
				}
				txSummaryTotal.Profit = txSummaryTotal.Income - txSummaryTotal.Fee
				txSummaryTotal.Balance = user.Balance
				txSummaryTotal.TodayIncome = int(total)
				return c.JSON(http.StatusOK, txSummaryTotal)
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

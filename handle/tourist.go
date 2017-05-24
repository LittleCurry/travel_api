package handle

import (
	"github.com/labstack/echo"
	//"git.iguiyu.com/park/api/vm"
	"net/http"
	//"git.iguiyu.com/park/struct/model"
	//"git.iguiyu.com/park/api/db"
	"github.com/curry/travel_api/db"
	"github.com/curry/travel_api/vm"
	"github.com/curry/travel_api/model"
	"github.com/jinzhu/copier"
	"strconv"
)

func TouristList(c echo.Context) (err error) {
	tourists := make([]model.Tourist, 0)
	err1 := db.MySQL().Find(&tourists)
	if err1 != nil {
		return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: err1.Error()})
	}
	txTourist := make([]vm.TxTourist, 0)
	copier.Copy(&txTourist, tourists)
	return c.JSON(http.StatusOK, txTourist)
}

func TouristPage(c echo.Context) (err error) {
	start, err1 := strconv.Atoi(c.Param("start"))
	if err1 == nil {
		limit, err2 := strconv.Atoi(c.Param("limit"))
		if err2 == nil {
			tourists := make([]model.Tourist, 0)
			err3 := db.MySQL().Limit(limit, start).Find(&tourists)
			if err3 == nil {
				txTourist := make([]vm.TxTourist, 0)
				copier.Copy(&txTourist, tourists)
				return c.JSON(http.StatusOK, txTourist)
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func DeleteTourist(c echo.Context) (err error) {
	touristId, err1 := strconv.Atoi(c.Param("id"))
	if err1 == nil {
		affected, err2 := db.MySQL().Id(touristId).Delete(&model.Tourist{})
		if err2 == nil && affected > 0 {
			return c.JSON(http.StatusOK, nil)
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

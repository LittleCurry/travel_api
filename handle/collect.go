package handle

import (
	"github.com/jinzhu/copier"
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"github.com/curry/travel_api/db"
	"github.com/curry/travel_api/model"
	"github.com/curry/travel_api/vm"
)

func CollectOrCancel(c echo.Context) (err error) {
	id, err1 := strconv.Atoi(c.Param("id"))
	if err1 == nil {
		//var affected int64
		var err2 error
		var err3 error
		//var has string

		switch c.Param("addDel") {
		case "add":
			_, err2 = db.MySQL().Id(id).Cols("collected").Update(&model.Tourist{Collected: 1})
			tourist := model.Tourist{}
			_, err3 := db.MySQL().Where("id=?", id).Get(&tourist)
			if err3 == nil {
				collect := model.Collect{}
				copier.Copy(&collect, &tourist)
				//collect.CreatDate = time.Now()
				_, err3 = db.MySQL().Insert(&collect)
			}
		case "del":
			_, err2 = db.MySQL().Id(id).Cols("collected").Update(&model.Tourist{Collected: 0})
			_, err3 = db.MySQL().Where("tourist_id=?", id).Delete(&model.Collect{})
			//_, err3 = db.MySQL().Id(id).Delete(&model.Collect{})
			//tourist := model.Tourist{}
			//_, err3 = db.MySQL().Where("id=?", id).Get(&tourist)
			//if err3 == nil {
				//collect := model.Collect{}
				//copier.Copy(&collect, &tourist)
				//collect.CreatDate = time.Now()
				//_, err3 = db.MySQL().Delete(&collect)
			//}
		}
		if err2 == nil && err3 == nil {
			return c.JSON(http.StatusOK, nil)
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func CollectList(c echo.Context) (err error) {
	start, err1 := strconv.Atoi(c.Param("start"))
	if err1 == nil {
		limit, err2 := strconv.Atoi(c.Param("limit"))
		if err2 == nil {
			collects := make([]model.Collect, 0)
			err3 := db.MySQL().Limit(limit, start).Find(&collects)
			if err3 == nil {
				txCollect := make([]vm.TxCollect, 0)
				copier.Copy(&txCollect, collects)
				return c.JSON(http.StatusOK, txCollect)
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func CountCollect(c echo.Context) (err error) {
	//userId := c.Get("userId").(string)
	total, err1 := db.MySQL().Count(&model.Collect{})
	if err1 !=nil {
		return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: err1.Error()})
	}
	return c.JSON(http.StatusOK, &vm.TxUnRead{Count: int(total)})
}

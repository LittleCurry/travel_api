package handle

import (
	"net/http"
	"strconv"

	"git.iguiyu.com/park/api/db"
	"git.iguiyu.com/park/api/vm"
	"git.iguiyu.com/park/struct/model"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
)

func ListMsg(c echo.Context) (err error) {
	start, err1 := strconv.Atoi(c.Param("start"))
	if err1 == nil {
		limit, err2 := strconv.Atoi(c.Param("limit"))
		if err2 == nil {
			userId := c.Get("userId").(string)
			msgs := make([]model.Message, 0)
			err3 := db.MySQL().Where("owner_user_id=?", userId).And("`deleted` = 0").Limit(limit, start).Desc("create_date").Find(&msgs)
			if err3 == nil {
				txMsgs := make([]vm.TxMessage, 0)
				copier.Copy(&txMsgs, &msgs)
				return c.JSON(http.StatusOK, txMsgs)
			}
		}
	}

	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func MarkRead(c echo.Context) (err error) {
	id, err1 := strconv.Atoi(c.Param("id"))
	if err1 == nil {
		userId := c.Get("userId").(string)
		affected, err2 := db.MySQL().Id(id).Where("owner_user_id=?", userId).Update(&model.Message{Read: 1})
		if err2 == nil && affected > 0 {
			return c.JSON(http.StatusOK, nil)
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func MarkAllToRead(c echo.Context) (err error) {
	userId := c.Get("userId").(string)
	_, err1 := db.MySQL().Where("owner_user_id=?", userId).Update(&model.Message{Read: 1})
	if err1 == nil {
		return c.JSON(http.StatusOK, nil)
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func MarkDelete(c echo.Context) (err error) {
	id, err1 := strconv.Atoi(c.Param("id"))
	if err1 == nil {
		userId := c.Get("userId").(string)
		affected, err2 := db.MySQL().Id(id).Where("owner_user_id=?", userId).Update(&model.Message{Deleted: 1})
		if err2 == nil && affected > 0 {
			return c.JSON(http.StatusOK, nil)
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func MarkAllToDeleted(c echo.Context) (err error) {
	userId := c.Get("userId").(string)
	_, err1 := db.MySQL().Where("owner_user_id=?", userId).Update(&model.Message{Deleted: 1})
	if err1 == nil {
		return c.JSON(http.StatusOK, nil)
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func CountUnRead(c echo.Context) (err error) {
	userId := c.Get("userId").(string)
	total, err1 := db.MySQL().Where("owner_user_id=?", userId).And("`read` = 0").Count(&model.Message{})
	if err1 == nil {
		return c.JSON(http.StatusOK, &vm.TxUnRead{Count: int(total)})
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

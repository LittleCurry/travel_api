package handle

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"git.iguiyu.com/park/api/db"
	"git.iguiyu.com/park/api/misc"
	"git.iguiyu.com/park/api/vm"
	"git.iguiyu.com/park/struct/model"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
)

const (
	OFFLINE = iota
	LOW_BATTERY
	GSM_HALT
)

/* filter=error,offline,lowbatt,gsmhalt */
func ListLocks(c echo.Context) (err error) {
	userId := c.Get("userId").(string)
	filter := c.QueryParam("filter")

	lockers := make([]model.Locker, 0)
	var err1 error
	switch filter {
	case "":
		err1 = db.MySQL().Where("owner_user_id=?", userId).And("active = 1").Find(&lockers)
	case "error":
		err1 = db.MySQL().Where("owner_user_id=?", userId).And("active = 1").And("error <> 0").Find(&lockers)
	case "off":
		err1 = db.MySQL().Where("owner_user_id=?", userId).And("active = 1").And("available = 0").Find(&lockers)
	case "lowbatt":
		err1 = db.MySQL().Where("owner_user_id=?", userId).And("active = 1").And("error & ? <> 0", 1<<LOW_BATTERY).Find(&lockers)
	case "gsmhalt":
		err1 = db.MySQL().Where("owner_user_id=?", userId).And("active = 1").And("error & ? <> 0", 1<<GSM_HALT).Find(&lockers)
	}

	if err1 == nil {
		txLocks := make([]vm.TxLockForList, 0)
		copier.Copy(&txLocks, lockers)
		return c.JSON(http.StatusOK, txLocks)
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func ShowLock(c echo.Context) (err error) {
	lockId, err1 := strconv.Atoi(c.Param("id"))
	if err1 == nil {
		userId := c.Get("userId").(string)
		locker := model.Locker{}
		has, err2 := db.MySQL().Id(lockId).Where("owner_user_id=?", userId).Get(&locker)
		if err2 == nil && has {
			profile := model.OwnerProfile{}
			has2, err3 := db.MySQL().Id(locker.OwnerProfileId).Get(&profile)
			if err3 == nil && has2 {
				txLock := vm.TxLockDetail{}
				copier.Copy(&txLock, &locker)
				txLock.Pname = profile.Pname
				return c.JSON(http.StatusOK, txLock)
			}

		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func GetLock(c echo.Context) (err error) {
	rxLocker := vm.RxLock{}
	err0 := c.Bind(&rxLocker)

	if err0 == nil {
		userId := c.Get("userId").(string)
		locker := model.Locker{}
		has, err1 := db.MySQL().Where("qrcode=?", rxLocker.QRCode).Get(&locker)
		//fmt.Println("err1:", err1)
		if err1 == nil && has {
			if locker.Active == 1 {
				if locker.OwnerUserId != misc.FormatUserIdToInt(userId) {
					return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "此地锁已激活"})
				} else {
					/*确实为本人的地锁，返回地锁的相关信息*/
					profile := model.OwnerProfile{}
					has, err2 := db.MySQL().Id(locker.OwnerProfileId).Get(&profile)

					if err2 == nil && has {
						txLocker := vm.TxLock{}
						copier.Copy(&txLocker, &locker)
						txLocker.Pname = profile.Pname
						return c.JSON(http.StatusOK, txLocker)
					} else {
						fmt.Println("err2:", err2)
					}
				}
			} else {
				txUnActivateLock := vm.TxUnActivatedLock{
					Id:      locker.Id,
					ShortId: locker.ShortId}
				return c.JSON(http.StatusOK, txUnActivateLock)
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func BindLock(c echo.Context) (err error) {
	lockId, err1 := strconv.Atoi(c.Param("id"))
	if err1 == nil {
		userId := c.Get("userId").(string)
		locker := model.Locker{}
		has, err2 := db.MySQL().Id(lockId).Get(&locker)
		if err2 == nil && has {
			//fmt.Println("owner:", locker.OwnerUserId)
			if locker.Active == 0 || locker.OwnerUserId == 0 || locker.OwnerUserId == misc.FormatUserIdToInt(userId) {
				rxLocker := vm.RxLockToBind{}
				//fmt.Println(c.Request().Body)
				err3 := c.Bind(&rxLocker)
				//fmt.Println("err3", err3)
				if err3 == nil {
					copier.Copy(&locker, &rxLocker)
					locker.OwnerUserId = misc.FormatUserIdToInt(userId)
					locker.Active = 1
					locker.ActiveDate = time.Now()
					locker.Available = 1
					_, err4 := db.MySQL().Id(lockId).Update(&locker)
					if err4 == nil {
						db.MySQL().Exec("UPDATE owner_user SET lockers_owned = (select Count(*) from locker where owner_user_id = ?) where id = ?", userId, userId)
						return c.JSON(http.StatusOK, nil)
					}
				}
			} else {
				return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "此地锁已绑定"})
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func TurnOnOff(c echo.Context) (err error) {
	lockId, err1 := strconv.Atoi(c.Param("id"))
	if err1 == nil {
		userId := c.Get("userId").(string)
		var affected int64
		var err2 error
		switch c.Param("onoff") {
		case "on":
			affected, err2 = db.MySQL().Id(lockId).Where("owner_user_id=?", userId).Cols("available").Update(&model.Locker{Available: 1})
		case "off":
			affected, err2 = db.MySQL().Id(lockId).Where("owner_user_id=?", userId).Cols("available").Update(&model.Locker{Available: 0})
		}
		if err2 == nil && affected > 0 {
			return c.JSON(http.StatusOK, nil)
		}

	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func ChangePlan(c echo.Context) (err error) {
	lockId, err1 := strconv.Atoi(c.Param("id"))
	if err1 == nil {
		planId, err2 := strconv.Atoi(c.Param("planid"))
		if err2 == nil {
			userId := c.Get("userId").(string)
			affected, err3 := db.MySQL().Id(lockId).Where("owner_user_id=?", userId).Cols("owner_profile_id").Update(&model.Locker{OwnerProfileId: planId})
			if err3 == nil && affected > 0 {
				return c.JSON(http.StatusOK, nil)
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func DownOrUp(c echo.Context) (err error) {
	lockId, err1 := strconv.Atoi(c.Param("id"))
	if err1 == nil {
		userId := c.Get("userId").(string)
		has, err2 := db.MySQL().Id(lockId).Where("owner_user_id=?", userId).Get(&model.Locker{})
		if err2 == nil && has {
			var err3 error
			switch c.Param("downup") {
			case "up":
				err3 = misc.UpLocker(lockId)
			case "down":
				err3 = misc.DownLocker(lockId)
			}
			if err3 == nil {
				return c.JSON(http.StatusOK, nil)
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

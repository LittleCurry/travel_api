package handle

import (
	"net/http"

	"time"

	"fmt"

	"git.iguiyu.com/park/api/db"
	"git.iguiyu.com/park/api/misc"
	"git.iguiyu.com/park/api/vm"
	"git.iguiyu.com/park/misc/globals"
	"git.iguiyu.com/park/struct/model"
	"github.com/labstack/echo"
)

func Login(c echo.Context) (err error) {
	rxUser := vm.RxUser{}
	if err1 := c.Bind(&rxUser); err1 == nil {
		user := model.OwnerUser{}
		has, err2 := db.MySQL().Where("loginname=?", rxUser.LoginName).And("passwd=?", misc.Md5Password(rxUser.Password)).Get(&user)
		if err2 == nil && has {
			txUser := &vm.TxUser{}
			txUser.AccessToken = globals.GenerateSession()
			db.Redis().Set("AccessToken-"+txUser.AccessToken, user.Id, 7*24*time.Hour)
			txUser.TokenType = "Bearer"

			user.LastLoginDate = time.Now()
			_, err3 := db.MySQL().Id(user.Id).Cols("last_login_date").Update(&user)
			if err3 == nil {
				return c.JSON(http.StatusOK, txUser)
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "帐号或密码错误"})
}

func SignUp(c echo.Context) (err error) {
	rxRegUser := vm.RxRegUser{}
	if err1 := c.Bind(&rxRegUser); err1 == nil {
		user := model.OwnerUser{}
		has, err2 := db.MySQL().Where("loginname=?", rxRegUser.Loginname).Limit(1, 0).Get(&user)
		if err2 == nil {
			if has {
				return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "登录名重复"})
			} else {
				code, err3 := db.Redis().Get(db.REDIS_CODE_PREFIX + rxRegUser.Loginname).Result()
				if err3 == nil {
					if code == rxRegUser.Code {
						fmt.Println("insert into DB")
						user.Loginname = rxRegUser.Loginname
						user.Nickname = rxRegUser.Nickname
						user.Passwd = misc.Md5Password(rxRegUser.Passwd)
						user.CreateDate = time.Now()
						user.Type = 1
						affected, err4 := db.MySQL().Insert(&user)
						if err4 == nil && affected > 0 {
							return c.JSON(http.StatusOK, nil)
						}
					} else {
						return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "验证码错误"})
					}
				} else {
					return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "验证码过期"})
				}
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "注册错误"})
}

func ResetPassword(c echo.Context) (err error) {
	rxRegUser := vm.RxNewPassword{}
	if err1 := c.Bind(&rxRegUser); err1 == nil {
		user := model.OwnerUser{Loginname: rxRegUser.Loginname}
		has, err2 := db.MySQL().Get(&user)
		if err2 == nil {
			if !has {
				return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "登录名不存在"})
			} else {
				code, err3 := db.Redis().Get(db.REDIS_CODE_PREFIX + rxRegUser.Loginname).Result()
				if err3 == nil {
					if code == rxRegUser.Code {
						user.Passwd = misc.Md5Password(rxRegUser.Passwd)
						affected, err4 := db.MySQL().Id(user.Id).Cols("passwd").Update(&user)
						if err4 == nil && affected > 0 {
							return c.JSON(http.StatusOK, nil)
						}
					} else {
						return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "验证码错误"})
					}
				}
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "密码重置错误"})
}

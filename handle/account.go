package handle

import (
	"net/http"

	"strconv"
	"time"

	"git.iguiyu.com/park/api/db"
	"git.iguiyu.com/park/api/joinmodel"
	"git.iguiyu.com/park/api/misc"
	"git.iguiyu.com/park/api/vm"
	"git.iguiyu.com/park/struct/model"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
)

/*
func ShowAccount(c echo.Context) (err error) {
	token := misc.GetAccessToken(c.Request().Header[echo.HeaderAuthorization][0])
	return c.JSON(http.StatusOK, token)
}

func UpdateAccount(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, nil)
}
*/

func AddBankAccount(c echo.Context) (err error) {
	rxBankAccount := vm.RxBankAccount{}
	if err1 := c.Bind(&rxBankAccount); err1 == nil {
		bank := model.Bank{}
		has, err2 := db.MySQL().Id(rxBankAccount.BankId).Get(&bank)
		if err2 == nil && has {
			userId := c.Get("userId").(string)
			account := model.Account{}
			copier.Copy(&account, &rxBankAccount)
			account.BankName = bank.BankName
			account.CreateTime = time.Now()
			account.OwnerUserId = misc.FormatUserIdToInt(userId)
			changed, err3 := db.MySQL().Insert(&account)
			if err3 == nil && changed > 0 {
				return c.JSON(http.StatusOK, nil)
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func ListBankAccount(c echo.Context) (err error) {
	userId := c.Get("userId").(string)
	bankaccounts := make([]joinmodel.AccountBank, 0)

	err1 := db.MySQL().Join(
		"INNER",
		"bank",
		"bank.id = account.bank_id").Where(
		"account.owner_user_id=?", userId).Find(&bankaccounts)
	//err1 := db.MySQL().Where("owner_user_id=?", userId).Find(&bankaccounts)
	if err1 == nil {
		txBankAccounts := make([]vm.TxBankAccount, 0)
		copier.Copy(&txBankAccounts, &bankaccounts)
		return c.JSON(http.StatusOK, txBankAccounts)
	}

	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func GetBankList(c echo.Context) (err error) {
	txBanks := make([]vm.TxBankType, 0)
	banks := make([]model.Bank, 0)
	err1 := db.MySQL().Find(&banks)
	if err1 == nil {
		copier.Copy(&txBanks, &banks)
		return c.JSON(http.StatusOK, txBanks)
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func Withdraw(c echo.Context) (err error) {
	rxWithdraw := vm.RxWithdraw{}
	if err1 := c.Bind(&rxWithdraw); err1 == nil {
		userId := c.Get("userId").(string)
		ownerUser := model.OwnerUser{}
		has, err2 := db.MySQL().Id(userId).Get(&ownerUser)
		if err2 == nil && has {
			code, err3 := db.Redis().Get(db.REDIS_CODE_PREFIX + ownerUser.Loginname).Result()
			if err3 == nil {
				if code == rxWithdraw.Code {
					withdraw := model.Withdraw{}
					copier.Copy(&withdraw, rxWithdraw)
					withdraw.CreateTime = time.Now()
					withdraw.OwnerUserId = misc.FormatUserIdToInt(userId)
					withdraw.Status = 1
					changed, err2 := db.MySQL().Insert(&withdraw)
					if err2 == nil && changed > 0 {
						return c.JSON(http.StatusOK, nil)
					}
				} else {
					return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "验证码错误"})
				}
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func DeleteBankAccount(c echo.Context) (err error) {
	accountid, err1 := strconv.Atoi(c.Param("id"))

	if err1 == nil {
		userId := c.Get("userId").(string)
		//widthdraws := make([]model.Withdraw, 0)
		total, err2 := db.MySQL().Where("account_id = ?", accountid).Count(&model.Withdraw{})
		if err2 == nil {
			if total == 0 {
				affected, err3 := db.MySQL().Id(accountid).Where("owner_user_id=?", userId).Delete(&model.Account{})
				if err3 == nil && affected > 0 {
					return c.JSON(http.StatusOK, nil)
				}
			} else {
				return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "此帐户已有提现记录，不能删除"})
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func ListWithdraw(c echo.Context) (err error) {
	start, err1 := strconv.Atoi(c.Param("start"))
	if err1 == nil {
		limit, err2 := strconv.Atoi(c.Param("limit"))
		if err2 == nil {
			userId := c.Get("userId").(string)
			widthdraws := make([]joinmodel.WithdrawAccountBank, 0)
			err3 := db.MySQL().Join(
				"INNER",
				"account",
				"account.id = withdraw.account_id").Join(
				"INNER",
				"bank",
				"bank.id = account.bank_id").Where(
				"account.owner_user_id=?", userId).Desc("withdraw.create_time").Limit(limit, start).Find(&widthdraws)
			if err3 == nil {
				txWithdraws := make([]vm.TxWithdraw, 0)
				copier.Copy(&txWithdraws, &widthdraws)
				return c.JSON(http.StatusOK, txWithdraws)
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func SendWithdrawCode(c echo.Context) (err error) {
	userId := c.Get("userId").(string)
	deviceId := c.Param("deviceid")

	ownerUser := model.OwnerUser{}
	has, err0 := db.MySQL().Id(userId).Get(&ownerUser)
	if err0 == nil && has {
		err1 := checkRequestPerDay(deviceId)
		if err1 == nil {
			err2 := sendVerifyCode(ownerUser.Loginname)
			if err2 == nil {
				return c.JSON(http.StatusOK, nil)
			}
		} else {
			return c.JSON(http.StatusBadRequest, err1.Error())
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "验证码发送失败"})
}

func RegDevice(c echo.Context) (err error) {
	userId := c.Get("userId").(string)
	deviceId := c.Param("deviceid")
	_, err1 := db.MySQL().Id(userId).Cols("device_id").Update(&model.OwnerUser{DeviceId: deviceId})
	if err1 == nil {
		return c.JSON(http.StatusOK, nil)
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

package handle

import (
	"errors"
	"net/http"
	"time"

	"git.iguiyu.com/park/api/db"
	"git.iguiyu.com/park/api/vm"
	"git.iguiyu.com/park/misc/globals"
	"git.iguiyu.com/park/struct/model"
	"github.com/denverdino/aliyungo/sms"
	"github.com/labstack/echo"
)

func smsCode(phone, code string) (err error) {
	ID := "LTAI1BBFOxU6Yjml"
	SECRET := "G0wrOThpwXFYTMTxXvF5HlKFBN5he4"
	SIGNAME := "鲑鱼出行"
	TEMPCODE := "SMS_53515144"
	client := sms.NewClient(ID, SECRET)
	return client.SendSms(&sms.SendSmsArgs{SignName: SIGNAME,
		TemplateCode: TEMPCODE,
		RecNum:       phone,
		ParamString:  `{"number": "` + code + `"}`})
}

func checkRequestPerDay(deviceId string) error {
	requestPerDay, _ := db.Redis().Get(db.REDIS_REQUEST_PER_DAY_OF_DEVICE + deviceId).Int64()
	switch {
	case requestPerDay == 0:
		db.Redis().Set(db.REDIS_REQUEST_PER_DAY_OF_DEVICE+deviceId, 1, 24*time.Hour)
	case requestPerDay > 3:
		return errors.New("每天验证码发送次数不能超过3次")
	default:
		db.Redis().Incr(db.REDIS_REQUEST_PER_DAY_OF_DEVICE + deviceId)
	}
	return nil
}

func sendVerifyCode(phoneNumber string) (err error) {
	randString := string(globals.Krand(4, globals.KC_RAND_KIND_NUM))
	db.Redis().Set(db.REDIS_CODE_PREFIX+phoneNumber, randString, 5*time.Minute)
	err = smsCode(phoneNumber, randString)
	return err
}

func SendCodeToNoneExistNumber(c echo.Context) (err error) {
	phoneNumber := c.Param("phonenumber")
	deviceId := c.Param("deviceid")

	has, err0 := db.MySQL().Where("loginname=?", phoneNumber).Limit(1, 0).Get(&model.OwnerUser{})
	if err0 == nil {
		if has {
			return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "此手机号码已被使用"})
		} else {
			err1 := checkRequestPerDay(deviceId)
			if err1 == nil {
				err2 := sendVerifyCode(phoneNumber)
				if err2 == nil {
					return c.JSON(http.StatusOK, nil)
				}
			} else {
				return c.JSON(http.StatusBadRequest, err1.Error())
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func SendCodeToExistNumber(c echo.Context) (err error) {
	phoneNumber := c.Param("phonenumber")
	deviceId := c.Param("deviceid")

	has, err0 := db.MySQL().Where("loginname=?", phoneNumber).Limit(1, 0).Get(&model.OwnerUser{})
	//fmt.Println("phonenumber:", phoneNumber)
	if err0 == nil {
		if !has {
			return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "手机号码尚未被注册过"})
		} else {
			err1 := checkRequestPerDay(deviceId)
			if err1 == nil {
				err2 := sendVerifyCode(phoneNumber)
				if err2 == nil {
					return c.JSON(http.StatusOK, nil)
				}
			} else {
				return c.JSON(http.StatusBadRequest, err1.Error())
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

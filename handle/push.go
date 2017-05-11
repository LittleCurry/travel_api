package handle

import (
	"github.com/denverdino/aliyungo/push"
	"fmt"
	"github.com/labstack/echo"
	"git.iguiyu.com/park/api/vm"
	"net/http"
	"git.iguiyu.com/park/struct/model"
	"git.iguiyu.com/park/api/db"
)

func Push(c echo.Context) (err error) {
	userId := c.Get("userId").(string)
	user := model.OwnerUser{}
	has, err1 := db.MySQL().Id(userId).Get(&user)
	txPush := vm.TxPush{}
	err2 := c.Bind(&txPush)
	if err1 == nil && has && err2 == nil {
		arg := push.PushArgs{
			AppKey:       23393293,
			PushType:     "NOTICE",
			DeviceType:   "ALL",
			Body:         txPush.Body,
			Target:       "ALL",
			TargetValue:  user.DeviceId,
			Title:        txPush.Title,
			StoreOffline: "true",
			//IOSApnsEnv:   "DEV",
		}
		client := push.NewClient("LTAI1BBFOxU6Yjml", "G0wrOThpwXFYTMTxXvF5HlKFBN5he4")
		res := client.Push(&arg)
		fmt.Println("result:", res)
		return c.JSON(http.StatusOK, nil)
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}


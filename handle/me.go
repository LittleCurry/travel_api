package handle

import (
	"net/http"

	"git.iguiyu.com/park/api/db"
	"git.iguiyu.com/park/api/misc"
	"git.iguiyu.com/park/api/vm"
	"git.iguiyu.com/park/struct/model"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
)

func Me(c echo.Context) (err error) {
	userId := c.Get("userId").(string)
	user := model.OwnerUser{Id: misc.FormatUserIdToInt(userId)}
	has, err2 := db.MySQL().Get(&user)
	if err2 == nil && has {
		txMe := new(vm.TxMe)
		copier.Copy(&txMe, &user)
		return c.JSON(http.StatusOK, txMe)
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "错误"})
}

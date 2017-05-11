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
)

func TouristList(c echo.Context) (err error) {
	tourists := make([]model.TouristSpot, 0)
	err1 := db.MySQL().Find(&tourists)
	if err1 != nil {
		return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: err1.Error()})
	}
	txTourist := make([]vm.TxTourist, 0)
	copier.Copy(&txTourist, tourists)
	return c.JSON(http.StatusOK, txTourist)
}

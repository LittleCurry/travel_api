package handle

import (
	"net/http"
	"github.com/jinzhu/copier"
	"github.com/curry/travel_api/db"
	"github.com/curry/travel_api/model"
	"github.com/labstack/echo"
	"github.com/curry/travel_api/vm"
	//"fmt"
	//"time"
	//"crypto/md5"
	//"io"
	//"os"
	"strconv"
	//"html/template"
	//"time"
)

func Share(c echo.Context) (err error) {
	start, err1 := strconv.Atoi(c.Param("start"))
	limit, err2 := strconv.Atoi(c.Param("limit"))
	lifes := make([]model.Life, 0)
	err3 := db.MySQL().Limit(limit, start).Find(&lifes)
	if err1 != nil || err2 != nil || err3 != nil{
		return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: err1.Error()})
	}
	txLifes := make([]vm.TxLife, 0)
	copier.Copy(&txLifes, lifes)
	return c.JSON(http.StatusOK, txLifes)

	//lifes := make([]model.Life, 0)
	//err1 := db.MySQL().Find(&lifes)
	//if err1 != nil {
	//	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: err1.Error()})
	//}
	//txLifes := make([]vm.TxLife, 0)
	//copier.Copy(&txLifes, lifes)
	//return c.JSON(http.StatusOK, txLifes)
}

func MakeShare(c echo.Context) (err error) {

	rxLife := model.Life{}
	if err1 := c.Bind(&rxLife); err1 != nil{
		return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: err1.Error()})
	}
	_, err2 := db.MySQL().Insert(&rxLife)
	if err2 != nil {
		return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: err2.Error()})
	}
	return c.JSON(http.StatusOK, nil)

/*


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
	*/
}

/*
func MakeShare(w http.ResponseWriter, r *http.Request, c echo.Context) (err error) {
	fmt.Println("method:", r.Method) //获取请求的方法


	tourists := make([]model.Tourist, 0)
	err1 := db.MySQL().Find(&tourists)
	// MakeShare函数 在post返回类型不一致


	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: err1.Error()})
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)  // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: err1.Error()})
		}
		defer f.Close()
		io.Copy(f, file)
	}

	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: err1.Error()})

}
*/

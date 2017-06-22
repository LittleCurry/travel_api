package handle

import (
	"net/http"
	"github.com/jinzhu/copier"
	"github.com/curry/travel_api/db"
	"github.com/curry/travel_api/model"
	"github.com/labstack/echo"
	"github.com/curry/travel_api/vm"
	"fmt"
	"time"
	"crypto/md5"
	"io"
	"os"
	"strconv"
	"html/template"
)

func Share(c echo.Context) (err error) {
	tourists := make([]model.Tourist, 0)
	err1 := db.MySQL().Find(&tourists)
	if err1 != nil {
		return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: err1.Error()})
	}
	txTourist := make([]vm.TxTourist, 0)
	copier.Copy(&txTourist, tourists)
	return c.JSON(http.StatusOK, txTourist)
}

func MakeShare(w http.ResponseWriter, r *http.Request, c echo.Context) (err error) {
	fmt.Println("method:", r.Method) //获取请求的方法


	tourists := make([]model.Tourist, 0)
	err1 := db.MySQL().Find(&tourists)


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

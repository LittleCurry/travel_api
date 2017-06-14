package handle

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"github.com/curry/travel_api/model"
	"github.com/curry/travel_api/db"
	"github.com/curry/travel_api/vm"
)

func Learn(c echo.Context) (err error) {
	// 实现
	total, err1 := db.MySQL().Count(&model.Collect{})
	if err1 !=nil {
		return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: err1.Error()})
	}
	return c.JSON(http.StatusOK, &vm.TxUnRead{Count: int(total)})
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)

	fmt.Printf( "最大值是 : %d\n", ret )
}


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

func swap(x, y string) (string, string) {
	return y, x
}

func use_array()  {
	/* 以下演示了数组初始化 */
	//var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 初始化数组中 {} 中的元素个数不能大于 [] 中的数字 */
	/* 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小 */
	//var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 该实例与上面的实例是一样的，虽然没有设置数组的大小 */
	//balance[4] = 50.0

	/* 数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。例如 */
	//float32 salary = balance[9]

	var n [10]int /* n 是一个长度为 10 的数组 */
	var i,j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j] )
	}
}

func use_ptr()  {
	var a int= 20   /* 声明实际变量 */
	var ip *int        /* 声明指针变量 */

	ip = &a  /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a  )

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )
}

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func use_struct()  {
	var Book1 Books        /* Declare Book1 of type Book */
	var Book2 Books        /* Declare Book2 of type Book */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook(&Book1)

	/* 打印 Book2 信息 */
	printBook(&Book2)
}

func use_slice(){
	var numbers = make([]int,3,5)
	printSlice(numbers)

	/////////////////
	var numbers2 []int
	printSlice(numbers2)

	if(numbers2 == nil){
		fmt.Printf("切片是空的")
	}
	/////////////////
	/* 创建切片 */
	numbers3 := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers3)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers3[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers3[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers3[4:])

	numbers4 := make([]int,0,5)
	printSlice(numbers4)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number5 := numbers[:2]
	printSlice(number5)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number6 := numbers[2:5]
	printSlice(number6)
}

func use_append()  {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2,3,4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1,numbers)
	printSlice(numbers1)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func printBook( book *Books ) {
	fmt.Printf( "Book title : %s\n", book.title);
	fmt.Printf( "Book author : %s\n", book.author);
	fmt.Printf( "Book subject : %s\n", book.subject);
	fmt.Printf( "Book book_id : %d\n", book.book_id);
}

func use_range()  {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func use_map() {
	var countryCapitalMap map[string]string
	/* 创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}

	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if(ok){
		fmt.Println("Capital of United States is", captial)
	}else {
		fmt.Println("Capital of United States is not present")
	}
}

func use_delete()  {
	/* 创建 map */
	countryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}

	fmt.Println("原始 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}

	/* 删除元素 */
	delete(countryCapitalMap,"France");
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}
}

func Factorial(x int) (result int) {
	if x == 0 {
		result = 1;
	} else {
		result = x * Factorial(x - 1);
	}
	return;
}

func use_digui() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(i))
}

func fibonaci(n int) int {
	if n < 2 {
		return n
	}
	return fibonaci(n-2) + fibonaci(n-1)
}

func digui2() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonaci(i))
	}
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func use_interface() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 	`error` 接口
func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

func use_error() {

	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

}

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)
	fmt.Printf( "最大值是 : %d\n", ret )

	c, d := swap("Mahesh", "Kumar")
	fmt.Println(c, d)
}






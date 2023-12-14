package main	// 1. package 进行包的声明, 建议: 包的声明这个包和所在的文件夹同名
// 2. main 包是程序的入口包, 一般 main 函数会放在这个包下

// 包名是从 $GOPATH/src/ 后开始计算的, 使用 / 进行路径分隔
import (
	"fmt"
	"utils/dbutils"
)

var Name string = "fly"

// var 声明定义一个变量，定义后，值可修改
func main() {
	var num int
	var num2 int = 30
	fmt.Println(num, num2)

	fmt.Println("你好，这是main函数的执行")
	dbutils.GetConn()
}

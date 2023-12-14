package main

import (
	"fmt"
	"unsafe"
)

// const 声明定义一个常量，定义后，值不可再修改
// := 定义一个变量，类型由后面的值决定。常用的定义方式
func main() {
	var num int
	var num1 = 20
	var num2 = 30
	num3 := 40
	var num4 int16 = 256

	fmt.Println(num, num1, num2, num3, num4)
	fmt.Println(unsafe.Sizeof(num))
	fmt.Printf("%T", num)
}

package main

import (
	"fmt"
)

// 总结：指针就是内存地址
func main() {
	a := 1
	fmt.Println(&a)
	var num *int = &a
	fmt.Println(*num)
	fmt.Println(num)

	var age int = 18
	fmt.Println(&age)

	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr本身这个存储空间的地址为：", &ptr)

	fmt.Println("--------- 通过指针改变指向值 -----------")
	var num2 int = 10
	fmt.Println(num2)

	var ptr2 *int = &num2
	*ptr2 = 20
	fmt.Println(num2)

	fmt.Println("--------- 指针变量接收的一定是地址值 ----------")
	var ptr3 *int = &num2
	fmt.Println(*ptr3)

	fmt.Println("-------- 指针变量的地址不可以不匹配 -----------")

	fmt.Println("-------- 基本数据类型（又叫值类型），都有对应的指针类型，形式为 *数据类型 --------")

}

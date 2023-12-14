package main

import "fmt"

func test(num int) {
	fmt.Println(num)
}

func test2(num int, testFunc func(int)) {
	fmt.Println("---------", num)
	testFunc(4)
}

// 支持对函数返回值命名
func test03(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}

func main() {
	a := test

	fmt.Printf("a的类型：%T, test函数的类型:%T\n", a, test)

	// 调用函数, test函数作为参数
	test2(10, test)

	sum, sub := test03(44, 22)
	fmt.Println(sum, sub)
}

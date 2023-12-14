package main

import (
	"fmt"
	"init/testutils"
)

// 1. 全局变量定义
var num int = test()

func test() int {
	fmt.Println("test函数被执行了")
	return 10
}

// 2. init 函数的调用
func init() {
	fmt.Println("init函数被执行了")
}

// 3. main 函数的调用
func main() {
	fmt.Println("main函数被执行了")
	fmt.Println(testutils.Age)

	// 定义匿名函数, 定义的同时调用
	result := func (num1 int, num2 int) int {
		return num1 + num2
	}(10, 20)
	fmt.Println(result)

	// 将匿名函数复制给一个变量, 这个变量实际就是函数类型的变量
	sub := func (num1 int, num2 int) int {
		return num1 + num2
	}
	result02 := sub(22, 33)
	fmt.Println(result02)
}
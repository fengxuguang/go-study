package main

import "fmt"

func main() {
	fmt.Println(add(30, 60))
}

func add(num1 int, num2 int) int {
	// 在 Golang 中, 程序遇到 defer 关键字, 不会立即执行 defer 后的语句, 而是将 defer 后的语句压入一个栈中, 然后继续执行函数后面的语句
	defer fmt.Println("num1=", num1)
	defer fmt.Println("num2=", num2)
	// 遇到 defer 关键字, 会将后面的代码语句压入栈中, 也会将相关的值同时拷贝入栈中, 不会随着函数后面的变化而变化
	num1 += 90
	num2 += 50

	// 栈的特点是：先进后出
	// 在函数执行完毕以后, 从栈中取出语句开始执行
	var sum int = num1 + num2
	fmt.Println("sum=", sum)
	return sum
}
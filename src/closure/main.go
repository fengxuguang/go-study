package main

import "fmt"

// 函数的名字: getSum 参数为空
// getSum 函数返回值为一个函数, 这个函数的参数是一个 int 类型的参数, 返回值也是 int 类型
func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum += num
		return sum
	}
}
// 闭包：返回的匿名函数 + 你们函数以外的变量 sum

func main() {
	f := getSum()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(getSum()(1))
	fmt.Println(getSum()(2))
}
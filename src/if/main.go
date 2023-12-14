package main

import "fmt"

// 全局变量 定义在函数外的变量
var name = "fly"

func main() {
	age := 18
	if age < 5 {
		fmt.Println("你的年龄为: ", age)
	} else if age < 20 {
		fmt.Println("你的年龄为：", age)
		if age > 15 {
			fmt.Println("你的你年龄大于 25")
		}
	} else {
		fmt.Println("你的年龄为: ", age)
	}
	fmt.Println("name=", name)
}

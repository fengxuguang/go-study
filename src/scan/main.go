package main

import "fmt"

func main() {
	// 实现功能：键盘录入学生的年龄、姓名、成绩、是否是 VIP
	// 方式一：ScanIn
	var age int
	fmt.Println("请录入学生的年龄：")
	// 传入 age 的地址的目的：在 ScanIn 函数中，对地址中的值进行改变的时候，实际外面的 age 被影响了
	fmt.Scanln(&age)

	var name string
	fmt.Println("请录入学生的姓名：")
	fmt.Scanln(&name)

	fmt.Printf("学生姓名为：%v,年龄为：%d", name, age)
}

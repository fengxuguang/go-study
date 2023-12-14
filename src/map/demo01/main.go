package main

import "fmt"

func main() {
	// 定义 map 变量
	var a map[int]string
	// 只声明 map 内存是没有分配空间
	// 必须通过 make 函数进行初始化, 才会分配空间
	a = make(map[int]string, 10)	// map 可以存放 10 个键值对
	// 将键值对存入 map 中
	a[2023001] = "张三"
	a[2023002] = "李四"
	a[2023003] = "王五"
	a[2023001] = "朱六"
	// 输出集合
	fmt.Println(a)
}
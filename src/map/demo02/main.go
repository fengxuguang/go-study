package main

import "fmt"

func main() {
	// 方式一：
	// 定义 map 变量
	var a map[int]string
	// 只声明 map 内存是没有分配空间
	// 必须通过 make 函数进行初始化, 才会分配空间
	a = make(map[int]string, 10)	// map 可以存放 10 个键值对
	// 将键值对存入 map 中
	a[2023001] = "张三"
	a[2023002] = "李四"
	// 输出集合
	fmt.Println(a)

	// 方式二：
	b := make(map[int]string)
	b[2023005] = "张三"
	b[2023006] = "李四"
	fmt.Println(b)

	// 方式三:
	c := map[int]string{
		2023007: "张三",
		2023008: "李四",
	}
	fmt.Println(c)
	value, flag := c[2023009]
	fmt.Println(value, flag)
}
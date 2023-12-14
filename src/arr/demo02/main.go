package main

import "fmt"

func main() {
	// 声明数组
	var arr [3]int8
	// 获取数组的长度
	fmt.Println(len(arr))

	// 打印数组
	fmt.Println(arr)

	// 证明 arr 中存储的地址值
	fmt.Printf("arr的地址为: %p\n", &arr)

	// 第一个空间的地址
	fmt.Printf("arr的地址: %p\n", &arr[0])

	// 第二个空间的地址
	fmt.Printf("arr的地址: %p\n", &arr[1])

	// 第三个空间的地址
	fmt.Printf("arr的地址: %p\n", &arr[2])

}
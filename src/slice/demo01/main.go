package main

import "fmt"

func main() {
	// 定义数组
	var intarr [6]int = [6]int{1, 2, 3, 4, 5, 6}

	// 切片构建在数组之上
	// 定义一个切片名字为 slice, []动态变化的数组长度不写, int 类型, intarr 是原数组
	var slice []int = intarr[1: 4]	// [) 左包右不包

	fmt.Println(slice)
	// 切片元素个数
	fmt.Println("slice的元素个数: ", len(slice))
	// 获取切片的容量; 容量可以动态变化
	fmt.Println("slice的容量:", cap(slice))
}
package main

import "fmt"

func main() {
	// 定义切片
	var a []int = []int{1, 4, 7, 3, 6, 9}
	// 再定义一个切片
	var b []int = make([]int, 10)

	// 拷贝, a里面的元素拷贝给b
	copy(b, a)
	fmt.Println(b)
}
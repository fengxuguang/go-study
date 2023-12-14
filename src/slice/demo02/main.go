package main

import "fmt"

func main() {
	// 定义切片: make 函数的三个参数: 1. 切片类型 2. 切片长度 3. 切片容量
	slice := make([]int, 4, 10)
	fmt.Println("切片的长度:", len(slice))
	fmt.Println("切片的容量:", cap(slice))

	slice[0] = 66
	slice[1] = 33
	fmt.Println(slice)
}
// ps: make 底层创建一个数组, 对外不可见, 所以不可以直接操作这个数组, 要通过 slice 去间接的访问各个元素, 不可以直接对数组进行维护/操作
package main

import "fmt"

func main() {
	// 定义数组
	var intarr [6]int = [6]int{1, 4, 7, 3, 6, 9}
	// 定义切片
	var slice []int = intarr[1: 4]
	fmt.Println(slice)

	slice2 := append(slice, 88, 50)
	fmt.Println(slice2)
	fmt.Println(slice)

	// 底层原理
	// 1. 底层追加元素的时候对数据进行扩容, 老数组扩容为新数组
	// 2. 创建一个新数组, 将老数组中的 4, 7, 3 复制到新数组中, 在新数组中追加 88, 50
	// 3. slice2 底层数组的指向  指向的是新数组
	// 4. 往往我们在使用追加的时候, 其实想要做的效果给 slice 追加
	slice = append(slice, 88, 50)
	fmt.Println(slice)

	slice3 := []int{99, 55}
	slice = append(slice, slice3...)
	fmt.Println(slice)
}
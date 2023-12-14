package main

import "fmt"

func main() {
	// 第一种
	var arr [3]int = [3]int{3, 4, 5}
	fmt.Println(arr)

	// 第二种
	var arr2 = [3]int{1, 2, 3}
	fmt.Println(arr2)

	// 第三种
	var arr3 = [...]int{2, 3, 4, 5, 6}
	fmt.Println(arr3)

	// 第四种
	var arr4 = [...]int{2:44, 0:22, 1:24, 3:55}
	fmt.Println(arr4)
}
type name interface {
	
}
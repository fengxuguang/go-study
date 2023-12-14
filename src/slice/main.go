package main

import (
	"fmt"
	"sort"
)

/**
 * 切片
 */
func main() {
	var s []int
	s1 := []int {1, 2, 3}
	
	fmt.Println(s)
	fmt.Println(s1)
	
	// 创建一个长度为 3、容量为 5 的整型切片
	slice := make([]int, 3, 5)
	// 创建一个长度为 3、容量为 3 的字符串切片
	slice2 := make([]string, 3)
	
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println("-------------")
	
	arr := [5]int{1, 2, 3, 4, 5}
	// 创建一个从 arr[1] 开始到 arr[4] 结束的切片
	slice3 := arr[1:4]
	// 创建一个新切片，容量等于原始切片的长度
	slice4 := slice3[:]
	slice3[2] = 6
	
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(arr)
	fmt.Println("------------")
	
	// 创建一个空的切片
	var slice6 []int
	// 向切片中追加元素
	slice6 = append(slice6, 4, 5, 6)
	fmt.Println(slice6)
	
	// 向切片中追加元素
	slice6 = append(slice6, 9)
	fmt.Println(slice6)
	
	println("------ for i ----")
	// for 循环遍历切片
	for i := 0; i < len(slice6); i ++ {
		println(slice6[i])
	}
	
	println("---------- for range ------")
	// for range 遍历切片
	for key, value := range slice6 {
		println(key, value)
	}
	
	println("------ copy ---------")
	// 创建一个包含 3 个元素的整数数组
	slice7 := make([]int, 2, 4)
	// 将 slice6 中的元素复制到 slice7 中
	copy(slice7, slice6)
	fmt.Println(slice7)
	
	println("---------- sort ----------")
	
	sortArr := []int{3, 5, 2, 4, 1}
	sort.Ints(sortArr)
	fmt.Println(sortArr)
}
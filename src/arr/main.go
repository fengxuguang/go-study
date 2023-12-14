package main

import "fmt"

/**
 * 数组
 */
func main() {
	var intArr [4]int
	var strArr [2]string
	var arr = [...]int{1, 2, 3, 4}
	var doubleArr = [...][2]int{{1, 2}, {3, 4}}
	
	fmt.Println(intArr)
	fmt.Println(strArr)
	fmt.Println(arr)
	fmt.Println(doubleArr)
}
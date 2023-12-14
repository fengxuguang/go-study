package main

import "fmt"

/**
 * 类型转换
 */
func main() {
	var num int = 42
	// 转浮点型
	var float float64 = float64(num)
	// 转整型无符号8位
	var ui uint8 = uint8(num)
	fmt.Println(num, float, ui)
	
	var n1 int32 = 12
	var n2 int64
//	var n3 int8
	
	n2 = int64(n1) + 20
	fmt.Println(n2)
}
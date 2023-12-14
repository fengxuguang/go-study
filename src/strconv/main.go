package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 23
	var n2 float32 = 3.14

	var s1 string = fmt.Sprintf("%d", num)
	fmt.Println(s1)

	var s2 string = fmt.Sprintf("%f", n2)
	fmt.Println(s2)

	// string -> bool
	var s3 string = "true"
	b, _ := strconv.ParseBool(s3)
	fmt.Printf("b的类型是：%T, b=%v\n", b, b)

	// string -> int64
	var s4 string = "23"
	var num1 int64
	num1, _ = strconv.ParseInt(s4, 10, 64)
	fmt.Printf("num1的类型是：%T, num1=%v\n", num1, num1)

	// stirng -> float64
	var s5 string = "3.14"
	var num2 float64
	num2, _ = strconv.ParseFloat(s5, 64)
	fmt.Printf("num2的类型是：%T, num2=%v\n", num2, num2)
}

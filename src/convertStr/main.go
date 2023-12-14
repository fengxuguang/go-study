package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 99
	var num2 float64 = 23.456
	var isTrue bool = true
	var str string = strconv.FormatInt(int64(num), 10)
	fmt.Println(str)
	
	str = strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Println(str)
	
	str = strconv.FormatBool(isTrue)
	fmt.Println(str)
}
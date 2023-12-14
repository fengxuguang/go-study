package main

import "fmt"

func cal(num1 int, num2 int) int {
	return num1 + num2
}

func cal2(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

func exchangeNum(num1 *int, num2 *int) {
	fmt.Println(*num1)
	fmt.Println(*num2)
	var t int
	t = *num1
	*num1 = *num2
	*num2 = t
}

func main() {
	sum := cal(10, 20)
	fmt.Println(sum)

	sum2, result := cal2(30, 10)
	fmt.Println(sum2)
	fmt.Println(result)

	fmt.Println("--------------")
	var num1 int = 10
	var num2 int = 20
	exchangeNum(&num1, &num2)
	fmt.Println(num1, num2)
}

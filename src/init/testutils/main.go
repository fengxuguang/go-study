package testutils

import "fmt"

var Age int

func init() {
	fmt.Println("testutils中的init函数被执行了")
	Age = 19
}
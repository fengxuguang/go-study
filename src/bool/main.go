package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i bool = true
	num := 3.14
	fmt.Printf("i=%T, 占用字节大小：%d", i, unsafe.Sizeof(i))
	fmt.Printf("\n%T", num)
}

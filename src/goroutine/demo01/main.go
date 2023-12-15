package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 1; i <= 5; i ++ {
		fmt.Println("test ---> hello golang, ", i)
		// 阻塞一秒
		time.Sleep(time.Second)
	}
}

func main() {
	go test()	// 开启协程

	for i := 1; i <= 10; i ++ {
		fmt.Println("main ---> hello golang, ", i)
		// 阻塞一秒
		time.Sleep(time.Second)
	}
}
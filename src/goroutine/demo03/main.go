package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 启动五个协程
	for i := 1; i <= 5; i ++ {
		// 协程开始的时候加 1
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("hello fly, ", n)
			// wg.Done()
		}(i)
	}

	// 主线程一直阻塞, 当 wg 减为 0了, 就停止
	wg.Wait()
}
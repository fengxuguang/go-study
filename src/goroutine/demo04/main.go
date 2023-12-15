package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
// 互斥锁
var lock sync.Mutex

var totalNum int = 0

func Sum() {
	defer wg.Done()
	for i := 1; i <= 100_000; i ++ {
		lock.Lock()
		totalNum += 1
		lock.Unlock()
	}
}

func Sub() {
	defer wg.Done()
	for i := 1; i <= 100_000; i ++ {
		lock.Lock()
		totalNum -= 1
		lock.Unlock()
	}
}

func main() {
	wg.Add(1)
	go Sum()

	wg.Add(1)
	go Sub()

	wg.Wait()
	fmt.Println(totalNum)
}
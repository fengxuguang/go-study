package main

import "fmt"

func main() {
	// 需要使用 make() 方法初始化 map
	var gomap = make(map[int]string)
	// 添加
	gomap[1] = string("hello")
	gomap[2]= string("golang")
	
	// 遍历
	for key,value := range gomap {
		fmt.Printf("index: %v, value: %v \n", key, value)
	}
	fmt.Println("---------------------")
	// 删除
	delete(gomap, 1)
	// 遍历
	for key,value := range gomap {
		fmt.Printf("index: %v, value: %v \n", key, value)
	}
}
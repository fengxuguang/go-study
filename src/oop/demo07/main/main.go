package main

import (
	"fmt"
	"oop/demo07/model"
)

func main() {
	// 跨包创建结构体 Student 实例
	s := model.Student{"fly", 18}
	fmt.Println(s)
}
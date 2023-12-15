package main

import "fmt"

type Student struct {
	Name string
	Age int
}

func main() {
	// 方式一：按照顺序赋值操作
	var s1 Student = Student{"fly", 23}
	fmt.Println(s1)

	// 方式二：按照指定类型
	var s2 Student = Student{
		Name: "fly2",
		Age: 22,
	}
	fmt.Println(s2)

	// 方式三：想要返回结构体的指针类型
	var s3 *Student = &Student{"fly3", 24}
	fmt.Println(*slice3)
}
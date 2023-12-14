package main

import "fmt"

// 定义老师结构体
type Teacher struct {
	Name string
	Age int
	School string
}

func main() {
	// 方式一
	var t1 Teacher
	t1.Name = "fly"
	t1.Age = 23
	t1.School = "清华大学"
	fmt.Println(t1)
	fmt.Println(t1.Age + 10)

	// 方式二
	var t2 Teacher = Teacher{"fly", 23, "中山大学"}
	fmt.Println(t2)

	// 方式三
	var t *Teacher = new(Teacher)
	(*t).Name = "fly"
	(*t).Age = 24
	t.School = "华南理工大学"
	fmt.Println(*t)

	// 方式四
	var t3 *Teacher = &Teacher{"fly", 23, "北京大学"}
	fmt.Println(*t3)
}
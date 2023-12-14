package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) test() {
	p.Name = "fly2"
	fmt.Println(p.Name)
}

func (p *Person) test2() {
	(*p).Name = "fly3"
	fmt.Println(p.Name)
}

func main() {
	// 创建结构体对象
	var p Person
	p.Name = "fly"
	p.test()
	fmt.Println(p.Name)

	fmt.Println("-------")
	(&p).test2()
	fmt.Println(p.Name)
}
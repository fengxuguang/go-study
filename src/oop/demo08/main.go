package main

import "fmt"

// 定义动物结构体
type Animal struct {
	Age int
	Weight float32
}

// 给 Animal 绑定方法, 喊叫
func (an *Animal) shout() {
	fmt.Println("我可以大声喊叫")
}

// 给 Animal 绑定方法, 自我展示
func (an *Animal) showInfo() {
	fmt.Printf("动物的年龄是：%v, 动物的体重是：%v\n", an.Age, an.Weight)
}

func (an *Cat) showInfo() {
	fmt.Printf("~~~~~动物的年龄是：%v, 动物的体重是：%v\n", an.Age, an.Weight)
}

// 定义结构体 Cat
type Cat struct {
	// 为了复用性, 体现继承思维, 加入匿名结构体
	Animal
}

// 对 Cat 绑定特有的方法
func (c Cat) scratch() {
	fmt.Println("我是小猫，我可以挠人")
}

func main() {
	// 创建 Cat 结构体示例
	cat := Cat{}
	// cat.Animal.Age = 3
	// cat.Animal.Weight = 10.6
	cat.Age = 3
	cat.Weight = 10.6

	cat.shout()
	cat.showInfo()
	cat.scratch()
}
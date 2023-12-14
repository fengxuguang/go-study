package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	u := user{"fly", 18}
	fmt.Println(u.name)
	
	// 结构体四种实例化方式
	// 键值对初始化
	user2 := &user {name: "fly2", age: 22}
	fmt.Println(user2.name)
	
	// 值列表
	user3 := &user{"fly3", 23}
	fmt.Println(user3)
	
	// new 关键字
	user4 := new(user)
	user4.name = "fly4"
	fmt.Println(user4.name)
	
	fmt.Println("--------")
	uA := userA {
		name: "fly",
		age: 18,
		addr: address {
			city: "广州",
		},
	}
	fmt.Println(uA)
	data,err := json.Marshal(uA)
	if err == nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	fmt.Println("------------")
	// json 转结构体
	var userA2 = &uA
	json.Unmarshal([]byte(data), userA2)
	fmt.Println(userA2)
}

type user struct {
	name string
	age int
}

type userA struct {
	name string
	age int
	addr address
}

type address struct {
	city string
}
package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	message, err := Hello("fly")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	test()
	fmt.Println("finish!")
	err2 := test2(10, 0)
	if err2 != nil {
		fmt.Println("自定义错误：", err2)
		panic(err2)
	}
	fmt.Println("正常执行下面流程...")
}

func test2(num1 int, num2 int) error {
	if num2 == 0 {
		return errors.New("除数不能为0")
	} else {
		result := num1 / num2
		fmt.Println(result)
		return nil
	}
}

func test() {
	// 利用 defer + recover 来捕获错误
	defer func() {
		// 调用 recover 内置函数, 可以捕获错误
		err := recover()
		// 如果没有捕获错误, 返回值为零值: nil
		if err != nil {
			fmt.Println("错误已经捕获")
			fmt.Println("err是：", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

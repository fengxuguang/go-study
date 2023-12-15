package main

import (
	"fmt"
	// "io/ioutil"
	"io"
	"os"
	"bufio"
)

func main() {
	// 打开文件
	file, err := os.Open("C:/Users/0422120001/Documents/go_test.txt")
	if err != nil {
		fmt.Println("文件打开失败")
	}

	// 当函数退出时, 让 file 关闭, 防止内存泄漏
	defer file.Close()

	// 创建一个流
	reader := bufio.NewReader(file)
	// 读取操作
	for {
		str, err2 := reader.ReadString('\n')
		fmt.Println(str)
		if err2 == io.EOF {
			break
		}	
	}
	
	// 结束
	fmt.Println("文件读取成功")
	
}
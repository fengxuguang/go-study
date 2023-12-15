package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// 打开文件
	file, err := os.OpenFile("C:/Users/0422120001/Documents/go_test.txt", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}

	// 关闭文件
	defer file.Close()

	// 写入文件操作
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i ++ {
		writer.WriteString("你好 飞\n")
	}

	// 流带缓冲区, 刷新数据
	writer.Flush()
}
package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	// 打开文件
	file, err := os.Open("C:/Users/0422120001/Documents/go_test.txt")

	if err != nil {
		fmt.Println("文件打开出错，对应错误为：", err)
	}
	
	// 没有出错, 输出文件
	fmt.Println("文件=%v", file)

	// 读取文件内容
	// 备注: 在下面的程序中不需要进行 Open\Close 操作, 因为文件的打开和关闭操作被封装在 ReadFile 函数内部了
	bytes, err3 := ioutil.ReadFile("C:/Users/0422120001/Documents/go_test.txt")
	if err3 != nil {
		fmt.Println("文件读取出错，错误为: ", err3)
		return
	}
	fmt.Println(bytes)
	fmt.Printf("%v\n", string(bytes))

	// 关闭文件
	err2 := file.Close()
	if err2 != nil {
		fmt.Println("文件关闭失败")
	} else {
		fmt.Println("文件关闭成功")
	}
}
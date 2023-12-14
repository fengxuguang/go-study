package dbutils

import "fmt"

// 函数名首字母大写, 函数才能被其他包访问
// 同一个包下, 不可以定义重名的函数
func GetConn() {
	fmt.Println("GetConn")
}
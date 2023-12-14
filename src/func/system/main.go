package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 1. 统计字符串的长度, 按字节进行统计
	str := "golang你好fly"
	fmt.Println(len(str))

	// 2. 方式一：对字符串进行遍历
	for i, value := range str {
		fmt.Printf("索引为：%d,具体的值为：%c\n", i, value)
	}

	// 2. 方式二：利用 r:=[]rune(str)
	r := []rune(str)
	for i := 0; i < len(r); i ++ {
		fmt.Printf("%c\n", r[i])
	}

	// 3. 字符串转整数
	num, _ := strconv.Atoi("666")
	fmt.Println(num)

	// 整数转字符串
	str1 := strconv.Itoa(666)
	fmt.Println(str1)

	// 4. 统计一个字符串有几个指定的子串
	count := strings.Count("golanggolang", "lan")
	fmt.Println(count)

	// 5. 不区分大小写的字符串比较
	flag := strings.EqualFold("golang", "GOLANG")
	fmt.Println(flag)

	// 6. 返回子串在字符串第一次出现的索引值, 如果没有返回 -1
	index := strings.Index("golang", "g")
	fmt.Println(index)

	// 7. 字符串替换
	str2 := strings.Replace("golang", "go", "java", 1)
	fmt.Println(str2)

}
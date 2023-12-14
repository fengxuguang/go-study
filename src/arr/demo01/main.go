package main

import "fmt"

func main() {
	// 实现的功能: 给出 5 个学生的成绩, 求出成绩的总和、平均数
	test1()
	arr()
}

func arr() {
	var score [5]int
	score[0] = 95
	score[1] = 91
	score[2] = 49
	score[3] = 66
	score[4] = 22

	var sum int = 0
	// 总和
	for i := 0; i < len(score); i ++ {
		sum += score[i]
	}
	avg := float32(sum) / float32(len(score))
	fmt.Println("总和：", sum)
	fmt.Println("平均分：", avg)
}

func test1() {
	score1 := 95
	score2 := 91
	score3 := 49
	score4 := 66
	score5 := 22

	// 求和
	sum := score1 + score2 + score3 + score4 + score5

	// 平均数
	avg := float32(sum) / 5
	fmt.Println("总和：", sum)
	fmt.Println("平均分：", avg)
}
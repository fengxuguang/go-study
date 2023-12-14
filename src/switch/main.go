package main

import "fmt"

func main() {
	score := 75
	switch {
		case score < 60:
			fmt.Println(score)
		case score < 80:
			fmt.Println(score)
		case score >= 80:
			fmt.Println(score)
	}
}
package main

import "fmt"

func main() {
	var word string = "hello \n world"
	var word2 string = `hello \n world`
	fmt.Println(word)
	fmt.Println(word2)
}
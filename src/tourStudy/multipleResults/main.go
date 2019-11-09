package main

import "fmt"

func swap(wordA, wordB string) (string, string) {
	return wordB, wordA
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

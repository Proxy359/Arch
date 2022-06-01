package main

import (
	"fmt"
)

func first(a, b int) int {
	return a + b
}

func second(s ...int) int {
	ss := 0
	for _, v := range s {
		ss = ss + v
	}
	return ss
}

func main() {
	fmt.Println(first(10, 20))
	fmt.Println(second(10, 20, 30, 40))
}
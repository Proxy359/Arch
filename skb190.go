package main

import (
	"fmt"
)

func first(a [8]int) [8]int {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
	return a
}

func main() {
	mas := [8]int{8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(first(mas))
}
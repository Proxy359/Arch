package main

import "fmt"

func main() {
	fmt.Println("Эта программа позволит возвести число в степень")
	fmt.Println("Введите сначала число")
	var sum int
	var x int
	fmt.Scan(&x)

	fmt.Println("Введите степень, в которую хотите возмести число")
	var y int
	fmt.Scan(&y)

	if y == 0 {
		fmt.Println("Ваш ответ: 1")
	}

	for z := 1; z <= y; z++ {
		sum += x * x
	}

	fmt.Println("Ваш ответ:", sum)
}

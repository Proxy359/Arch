package main

import "fmt"

func main() {
	fmt.Println("Эта программа позволит сложить 4 числа")

	for x := 1; x <= 4; x++ {
		fmt.Println("Введи", x, "-е число")
		sum := 0
		var input int
		fmt.Scan(&input)
		for i := 1; i <= 4; i++ {
			sum = sum + input
		}

		if x == 4 {
			fmt.Println("Итак, полученная сумма это", sum)
			return
		}

	}
}
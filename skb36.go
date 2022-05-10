package main

import "fmt"

func main() {

	fmt.Print("Эта программа позволяет понять, какой у вас выпал билет.", "\n", "\n",
		"Введите последовательно 4 числа билета:", "\n")
	var num1 int
	var num2 int
	var num3 int
	var num4 int
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	fmt.Scan(&num3)
	fmt.Scan(&num4)

	if num1 == num4 &&
		num3 == num3 {
		fmt.Print("\n", num1, num2, num3, num4, " -> зеркальный билет")
	} else if num1 == num3+1 && num2 == num1+1 && num4 == num2+1 {
		fmt.Print("\n", num1, num2, num3, num4, " -> счастливый билет")
	} else {
		fmt.Print("\n", num1, num2, num3, num4, " -> обычный билет")
	}
}

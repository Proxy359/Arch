package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Эта прога показывает, сколько элементов массива идет после искомого числа.")
	fmt.Println("Или возвраащет 0, если число не было найдено.", "\n", "\n")

	x := 0
	var y int
	y = rand.Intn(10)
	var mass [10]int
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(mass); i++ {
		mass[i] = rand.Intn(10)
		if mass[i] == y {
			x = len(mass) - i
			break
		}
	}
	fmt.Println(mass, "- это массив")
	fmt.Println(y, "- это искомый элемент")

	if x == 0 {
		fmt.Println(0, "элемент в массиве не обнаружен")
	} else {
		fmt.Println(x-1, "- сколько элементов идут после него")
	}
}
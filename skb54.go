package main

import "fmt"

func main() {

	fmt.Println("Эта программа поочередно заполняет 3 корзины яблоками.")

	fmt.Println("Введите кол-во яблок для 1 порзины.")
	var basket1 int
	fmt.Scan(&basket1)

	fmt.Println("Введите кол-во яблок для 2 порзины.")
	var basket2 int
	fmt.Scan(&basket2)

	fmt.Println("Введите кол-во яблок для 3 порзины.")
	var basket3 int
	fmt.Scan(&basket3)

	for {

		x1 := 1
		x1 = x1 + 1
		if x1 == basket1 {
			fmt.Println("Первая корзина наполнена")
			x1 = 0
		}

		x2 := 1
		x2 = x2 + 1
		if x2 == basket2 {
			fmt.Println("Вторая корзина наполнена")
			x2 = 0
		}

		x3 := 1
		x3 = x3 + 1
		if x3 == basket3 {
			fmt.Println("Тертья корзина наполнена")
			x3 = 0
		}

	}
}
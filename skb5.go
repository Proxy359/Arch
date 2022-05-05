package main

import "fmt"

func main () {
	var num int
		fmt.Println("Введи число, квадрат которого хочешь получить")
		fmt.Scan(&num)
		fmt.Println("Квадратом числа", num, "является число",num*num)
}
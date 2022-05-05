package main

import "fmt"

func main () {
	fmt.Println("Привет, я твой виртуальный помощник")
	fmt.Println("Чтобы пройти регистрацию, заполни пару строкчек:")
	
	fmt.Println("- введи свой username")
	var username string
	fmt.Scan(&username)

	fmt.Println("- введи пароль до 20 символов")
	var password string
	fmt.Scan(&password)

	fmt.Println("- уточни свой возраст")
	var userAge int
	fmt.Scan(&userAge)

	fmt.Println("Поздравляю,", username, "теперь вы зарегистрированы!")
	fmt.Println("Ваш пароль:", password)
	fmt.Println("Ваш возраст:", userAge)
}
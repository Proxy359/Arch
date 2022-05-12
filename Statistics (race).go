package main

import "fmt"

func main() {
	var lapNumber int = 4
	var engineBuff int = 254
	var wheelsBuff int = 93
	var rudderBuff int = 49
	var windDebuff int = -21
	var rainDebuff int = -17
	var driverSpeed int = engineBuff + wheelsBuff + rudderBuff + windDebuff + rainDebuff
	fmt.Println("===================")
	fmt.Println("Супер гонки. Круг", lapNumber)
	fmt.Println("===================")
	fmt.Print("Шумахер (", driverSpeed, ")\n")
	fmt.Println("===================")
	fmt.Println("Водитель: Шумахер")
	fmt.Println("Скорость:", driverSpeed)
	fmt.Println("-------------------")
	fmt.Println("Оснащение")
	fmt.Print("Двигатель: +", engineBuff, "\n")
	fmt.Print("Колеса: +", wheelsBuff, "\n")
	fmt.Print("Руль: +", rudderBuff, "\n")
	fmt.Println("-------------------")
	fmt.Println("Действия плохой погоды")
	fmt.Println("Ветер:", windDebuff)
	fmt.Println("Дождь:", rainDebuff)
}

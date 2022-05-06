package main

import "fmt"

func main() {

	fmt.Print("Эта программа позволяет выяснить, сколько вы заплатите налогов.", "\n", "\n",
		"Введите сумму вашего дохода:", "\n")
	var revenue int
	fmt.Scan(&revenue)

	bigLimit := 50000
	smallLimit := 10000

	if revenue > bigLimit {
		fixTax2BigLimit := (bigLimit - smallLimit) / 10 * 2
		bigTax := (revenue - bigLimit) / 10 * 3
		fmt.Print("Ваш налог на доход составляет ", bigTax+fixTax2BigLimit, " рублей.")
	} else if revenue > smallLimit {
		smallTax := (revenue - smallLimit) / 10 * 2
		fmt.Print("Ваш налог на доход составляет ", smallTax, " рублей.")
	} else {
		fmt.Print("Вам не надо платить налогов.")
	}
}

package main

import "fmt"

func main() {

	fmt.Print("Эта программа позволяет выяснить, какие из 3 чисел больше 5,", "\n", "\n",
		"Введите первое число:", "\n")

	var resultFirst int
	fmt.Scan(&resultFirst)

	var resultSecond int
	fmt.Print("Введите второе число:", "\n")
	fmt.Scan(&resultSecond)

	var resultThird int
	fmt.Print("Введите третье число:", "\n")
	fmt.Scan(&resultThird)

	lowerBarrier := 5

	if resultFirst <= lowerBarrier &&
		resultSecond <= lowerBarrier &&
		resultThird <= lowerBarrier {
		fmt.Print("\n", "Чисел больше 5 не обнаружено:", "\n", "\n")
	} else {
		fmt.Print("\n", "Обнаружены следующие числа больше 5:", "\n")
	}

	if resultFirst > lowerBarrier {
		fmt.Print(resultFirst, "\n")
	}

	if resultSecond > lowerBarrier {
		fmt.Print(resultSecond, "\n")
	}

	if resultThird > lowerBarrier {
		fmt.Print(resultThird, "\n")
	}
}
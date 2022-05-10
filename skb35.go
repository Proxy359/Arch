package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Print("Эта программа позволяет решить квадратное уравнение.", "\n", "\n",
		"Введите первый коэффициент:", "\n")
	var a float64
	fmt.Scan(&a)

	fmt.Print("Введите второй коэффициент:", "\n")
	var b float64
	fmt.Scan(&b)

	fmt.Print("Введите третий коэффициент:", "\n")
	var c float64
	fmt.Scan(&c)

	quartartB := math.Pow(b, 2)

	d := quartartB - 4*a*c

	if d > 0 {
		x1 := (-b + math.Sqrt(d)) / 2 * a
		x2 := (-b - math.Sqrt(d)) / 2 * a
		fmt.Print("\n", "\n", "Решение задачи: x1 = ", x1, " x2 = ", x2)
	} else if d == 0 {
		x := -b / 2 * a
		fmt.Print("\n", "\n", "Решение задачи: x = ", x)
	} else {
		fmt.Print("\n", "\n", "У задачи нет решения")
	}
}

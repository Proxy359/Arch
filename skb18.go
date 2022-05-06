package main

import (
	"fmt"
)

func main() {
	fmt.Print("Эта программа позволяет вычислять минимальное натуральное число из двух.","\n")
  
  var numA int 
  fmt.Print("Введи первое число.","\n")
fmt.Scan(&numA)

var numB int 
  fmt.Print("Введи второе число.","\n")
fmt.Scan(&numB)

if numA==numB {
  fmt.Print("Числа ",numA," и ",numB," равны.")
} else if numA>numB {
  fmt.Print("Миниматльное число это ",numB)
 } else {
  fmt.Print("Миниматльное число это ",numA)
 }
  }
package main

import (
	"fmt"
)

func main() {
	fmt.Print("Эта программа позволяет вычислять модуль из любого натурального числа.","\n","Напиши число и получишь модуль через 2-3 секунды!","\n")
  var num int 
fmt.Scan(&num)

  if num <0 {
    num=num*-1
  }

  fmt.Print("\n","Модуль этого числа равен ",num)
  }

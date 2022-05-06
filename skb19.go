package main

import (
	"fmt"
)

func main() {
	fmt.Print("Эта программа позволяет вычеслить самого умного астонавта.","\n")
  
  var ast1 int 
  fmt.Print("Введи IQ первого астронавта.","\n")
fmt.Scan(&ast1)

var ast2 int 
  fmt.Print("Введи IQ второго астронавта.","\n")
fmt.Scan(&ast2)

var ast3 int 
  fmt.Print("Введи IQ третьего астронавта.","\n")
fmt.Scan(&ast3)

if ast1>ast2 && ast1>ast3 {
  fmt.Print("Первый астронавт лучше подойдет на должность командира.")
} else if  ast2>ast1 && ast2>ast3 {
  fmt.Print("Второй астронавт лучше подойдет на должность командира.")
} else if  ast3>ast1 && ast3>ast2 {
  fmt.Print("Третий астронавт лучше подойдет на должность командира.")
} else {
  fmt.Print("Этим образом выбрать командира нельзя т.к. показатели IQ двух астронавтов равны.")
}
  }
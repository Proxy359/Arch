package main

import "fmt"

func main() {

	fmt.Print("Давай сыграем в угадай число (от 1 до 10).","\n","Загадал? Отлично!","\n",
            "Предположу, что это число 5!","\n", "\n",)

  min:=1
  max:=10
  num:=(max+min)/2
  fmt.Print("Если я угадал, напиши - да","\n",
            "Если не угадал, напили больше или меньше ")
  
var answer string
  fmt.Scan(&answer)
  
if answer == "да" {
    fmt.Print("Еее, мы победили!")
  } else if answer == "больше" {
  num=(max-min+1)/2
    fmt.Print("Ты загадал число ",num,"? ")
  } else {
   num=(max-min-1)/2
    fmt.Print("Ты загадал число ",num,"? ")
  }

var answerD string
  fmt.Scan(&answerD)
  
  
if answerD == "да" {
    fmt.Print("Еее, мы победили!")
  } else if answerD == "больше" {
 num=(max-min+1)/2
    fmt.Print("Ты загадал число ",num,"? ")
  } else {
     num=(max-min-1)/2
    fmt.Print("Ты загадал число ",num,"? ")
  }


var answerT string
  fmt.Scan(&answerT)
  
  
  
if answerT == "да" {
    fmt.Print("Еее, мы победили!")
  } else if answerT == "больше" {
  num=(max-min+1)/2
    fmt.Print("Ты загадал число ",num," ?")
  } else {
     num=(max-min-1)/2
    fmt.Print("Ты загадал число ",num," ?")
 
  }

}
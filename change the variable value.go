package main

import (
	"fmt"
)

func main () {
fmt.Println("Этот код позволит поменять значения двух переменных","\n")
  
  var a int
  fmt.Println("Введите желаемое значение а")
  fmt.Scan(&a)

  var b int
   fmt.Println("Введите желаемое значение b")
  fmt.Scan(&b)
  fmt.Println("Получается, что было:","\n","a=",a,"\n","b=",b,"\n")

  a=a+b
  b=a-b
  a=a-b

  fmt.Println("Но произошла магия и теперь:","\n","a=",a,"\n","b=",b,"\n")
}

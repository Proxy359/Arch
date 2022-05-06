package main

import (
	"fmt"
)

func main() {
	fmt.Print("Лор: вы выехали в соседний город.","\n","у вас есть 2 часа, чтобы доехать до пункта назвачения.","\n","С какой средней скоростью вы будете ехать? (в км.)","\n")
  var speed int 
fmt.Scan(&speed)

  traveledDistance:=2*speed
  fmt.Print("Таким образом, вы проехали ",traveledDistance," километров.")
  if traveledDistance >=200 {
    fmt.Print("\n","Вы приехали!")
  }
  }

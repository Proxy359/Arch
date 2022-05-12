package main

import (
	"fmt"
)

func main() {
	fmt.Print("Давайте посчитаем стоимость всех товаров.","\n")
  itemAmount:=0

  var firstPrise int
  fmt.Print("Уточните стоимость первого товара.","\n")
  fmt.Scan(&firstPrise)

  var secondPrise int
  fmt.Print("Уточните стоимость второго товара.","\n")
  fmt.Scan(&secondPrise)

  var thirdPrise int
  fmt.Print("Уточните стоимость третьего товара.","\n")
  fmt.Scan(&thirdPrise)
   itemAmount=thirdPrise+secondPrise+firstPrise

  discountBoundary:=10000
  if itemAmount > discountBoundary {
    fmt.Print("\n","Ого, к вашему заказу применима скидка 10%!")
    itemAmount=itemAmount/10*9
  }

  fmt.Print("\n","Итого, сумма ваших покупок составляет ", itemAmount," рублей.")
  }

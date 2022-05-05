package main

import "fmt"

func main() {
	var initialValue int
	var deliveryFee int
	var discount int
	fmt.Println("Калькулятор стоимости товара со скидкой.")
	fmt.Println("Стоимость товара:")
  fmt.Scan(&initialValue)
	fmt.Println("Стоимость доставки:")
  fmt.Scan(&deliveryFee)
	fmt.Println("Размер скидки:")
  fmt.Scan(&discount)
	fmt.Println("---------")
	fmt.Println("Итого:", initialValue + deliveryFee - discount)
}

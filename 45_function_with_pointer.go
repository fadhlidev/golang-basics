package main

import "fmt"

type Order struct {
	Total int
}

func IncrementTotalOrder(order *Order) {
	order.Total++
}

func main() {
	o := Order{0}
	fmt.Println(o)

	IncrementTotalOrder(&o)
	fmt.Println(o)
}

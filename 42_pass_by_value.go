package main

import "fmt"

type Address struct {
	City string
}

func main() {
	address1 := Address{"Solo"}
	address2 := address1

	address2.City = "Jogja"

	fmt.Println(address1) // Didn't change
	fmt.Println(address2)
}

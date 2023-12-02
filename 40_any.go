package main

import "fmt"

func main() {
	// typ any = interface {}
	var anything any

	anything = 1
	fmt.Println(anything)

	anything = "one"
	fmt.Println(anything)

	anything = true
	fmt.Println(anything)
}

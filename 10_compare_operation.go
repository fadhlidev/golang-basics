package main

import "fmt"

func main() {
	var isLower = 5 < 10
	var isGreater = 10 > 5
	var isEqualOrLower = 5 <= 10
	var isEqualOrGreater = 10 >= 5
	var isEqual = 10 == 10
	var isNotEqual = 5 != 10

	fmt.Println(isLower)
	fmt.Println(isGreater)
	fmt.Println(isEqualOrLower)
	fmt.Println(isEqualOrGreater)
	fmt.Println(isEqual)
	fmt.Println(isNotEqual)
}

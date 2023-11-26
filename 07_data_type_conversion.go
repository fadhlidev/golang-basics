package main

import "fmt"

func main() {
	var number int32 = 32459
	var toByte = byte(number)

	fmt.Println(number)
	fmt.Println(toByte)

	var text = "amazing"
	var z = text[3]
	fmt.Println(z)
	fmt.Println(string(z))
}

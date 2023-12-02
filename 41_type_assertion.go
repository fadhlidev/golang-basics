package main

import "fmt"

func main() {
	var value any
	value = "Fadhli Syaifullah"

	var name string
	name = value.(string)

	fmt.Println(name)
}

package main

import "fmt"

func main() {
	var greeting string
	var name = "Fadhli Syaifullah"
	var (
		firstName = "Fadhli"
		lastName  = "Syaifullah"
	)
	age := 17

	fmt.Println("My name is", name, "and im", age, "yo")

	greeting = "Hi there!"
	age = 18
	fmt.Println(greeting, "My name is", firstName, lastName, "and im", age, "yo now")
}

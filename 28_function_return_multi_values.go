package main

import "fmt"

func swapName(firstName, lastName string) (string, string) {
	return lastName, firstName
}

func main() {
	lastName, firstName := swapName("Fadhli", "Syaifullah")
	fmt.Println(lastName, firstName)

	_, name := swapName("Fadhli", "Syaifullah")
	fmt.Println(name)
}

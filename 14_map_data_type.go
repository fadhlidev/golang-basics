package main

import "fmt"

func main() {
	person := map[string]string{
		"firstName": "Fadhli",
		"lastName":  "Syaifullah",
		"country":   "indonesia",
	}

	fmt.Println(person)
	fmt.Println(person["firstName"])
	fmt.Println(person["country"])

	person["country"] = "Indonesia"
	println(person["country"])

	fmt.Println(len(person))
	delete(person, "country")
	fmt.Println(len(person))
	fmt.Println(person["country"])
	fmt.Println(person)
}

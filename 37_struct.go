package main

import "fmt"

type Person struct {
	FirstName, LastName string
	Age                 int
}

func main() {
	p := Person{
		FirstName: "Fadhli",
		LastName:  "syaifullah",
		Age:       17,
	}

	fmt.Println(p)

	p.Age++
	fmt.Println(p)

	p.LastName = "Syaifullah"
	fmt.Println(p)
}

package main

import (
	"fmt"
	"strings"
)

func splitName(fullName string) (firstName, middleName, lastName string) {
	names := strings.Split(fullName, " ")
	firstName = names[0]
	if len(names) == 2 {
		lastName = names[1]
	} else {
		middleName = names[1]
		lastName = names[2]
	}

	return firstName, middleName, lastName
}

func main() {
	firstName, _, lastName := splitName("Fadhli Syaifullah")
	fmt.Println(firstName)
	fmt.Println(lastName)
}

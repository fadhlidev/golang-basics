package main

import "fmt"

func main() {
	username := "boy"
	length := len(username)

	switch {
	case length > 10:
		fmt.Println("Username too long")
	case length < 5:
		fmt.Println("Username too short")
	default:
		fmt.Println("Valid username")
	}

}

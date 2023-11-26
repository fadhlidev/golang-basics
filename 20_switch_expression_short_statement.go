package main

import "fmt"

func main() {
	username := "iamagoodboy"

	switch length := len(username); length > 10 {
	case true:
		fmt.Println("Username too long")
	case false:
		fmt.Println("Valid username")
	}
}

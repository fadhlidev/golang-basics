package main

import "fmt"

func main() {
	username := "iamagoodboy"

	if length := len(username); length > 10 {
		fmt.Println("Username too long")
	}
}

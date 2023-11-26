package main

import "fmt"

func main() {
	condition := true
	if condition {
		fmt.Println("Run when condition is true (1)")
	}

	condition = false
	if condition {
		fmt.Println("Run when condition is true (2)")
	}
}

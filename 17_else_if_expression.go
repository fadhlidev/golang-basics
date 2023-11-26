package main

import "fmt"

func main() {
	condition1 := false
	condition2 := true
	if condition1 {
		fmt.Println("Run when condition 1 is true")
	} else if condition2 {
		fmt.Println("Run when condition 2 is true")
	} else {
		fmt.Println("Run when neither condition 1 or 2 is true")
	}
}

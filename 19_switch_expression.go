package main

import "fmt"

func main() {
	grade := "B"

	switch grade {
	case "A":
		fmt.Println("Great!")
	case "B":
		fmt.Println("Good")
	case "C":
		fmt.Println("Bad")
	default:
		fmt.Println("WTF!")
	}
}

package main

import "fmt"

func main() {
	defer func() {
		r := recover()

		if r != nil {
			fmt.Println("recover", r)
		}

		fmt.Println("Run after function main")
	}()

	panic("Ups!")

	fmt.Println("Run")
}

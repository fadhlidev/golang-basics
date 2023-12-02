package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Run after function main")
	}()

	fmt.Println("Run")
}

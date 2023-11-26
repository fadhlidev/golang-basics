package main

import "fmt"

func main() {
	counter := 0

	for counter < 10 {
		fmt.Println("Iteration", counter)
		counter++
	}

	fmt.Println("----------")

	for i := 10; i < 20; i++ {
		fmt.Println("Iteration", i)
	}
}

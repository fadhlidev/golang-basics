package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i == 10 {
			break
		}

		if i%2 == 0 {
			fmt.Println("Even Iteration", i)
			continue
		}
	}
}

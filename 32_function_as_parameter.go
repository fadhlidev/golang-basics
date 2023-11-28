package main

import "fmt"

func callMe(name string, callback func()) {
	fmt.Println("Hello", name)
	callback()
}

func main() {
	callMe("Fadhli", func() {
		fmt.Println("Hi, nice to meet you")
	})
}

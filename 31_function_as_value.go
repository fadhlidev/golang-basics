package main

import "fmt"

func bye() {
	fmt.Println("Bye!")
}

func main() {
	goodbye := bye
	goodbye()
}

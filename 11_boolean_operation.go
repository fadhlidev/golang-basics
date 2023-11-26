package main

import "fmt"

func main() {
	var TRUE = true
	var FALSE = false

	var AND = TRUE && FALSE
	fmt.Println(AND)

	var OR = TRUE || FALSE
	fmt.Println(OR)

	var NEGATION = !TRUE
	fmt.Println(NEGATION)
}

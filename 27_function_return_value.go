package main

import "fmt"

func sum(a int64, b int64) int64 {
	return a + b
}

func main() {
	result := sum(10, 20)
	fmt.Println(result)
}

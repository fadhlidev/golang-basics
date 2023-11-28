package main

import "fmt"

func sumAll(numbers ...int64) int64 {
	var total int64 = 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	result := sumAll(1, 2, 3, 4, 5)
	fmt.Println(result)

	numbers := []int64{1, 2, 3, 4, 5}
	fmt.Println(sumAll(numbers...))
}

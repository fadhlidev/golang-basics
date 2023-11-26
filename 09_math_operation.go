package main

import "fmt"

func main() {
	var (
		a = 100
		b = 5
	)

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	var c = 10
	c += 10
	fmt.Println(c)

	c -= 5
	fmt.Println(c)

	c *= 10
	fmt.Println(c)

	c /= 5
	fmt.Println(c)

	c %= 10
	fmt.Println(c)

	var d = 10
	d++
	fmt.Println(d)

	d--
	d--
	fmt.Println(d)
}

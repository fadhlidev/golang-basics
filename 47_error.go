package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divider can't be 0")
	}

	return a / b, nil
}
func main() {
	d1, err := divide(10, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d1)
	}

	d2, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d2)
	}
}

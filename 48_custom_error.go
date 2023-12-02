package main

import "fmt"

type NotFoundError struct {
	Message string
}

func (n NotFoundError) Error() string {
	return n.Message
}

func saveData(id int) error {
	if id == 0 {
		return NotFoundError{"Id 0 not found"}
	}

	return nil
}

func main() {
	err := saveData(1)
	if err != nil {
		fmt.Println(err)
	}

	err = saveData(0)
	if err != nil {
		fmt.Println(err)
	}
}

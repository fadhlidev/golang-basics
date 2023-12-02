package main

import "fmt"

type Teacher struct {
	Name string
}

func main() {
	teacher1 := new(Teacher)
	teacher2 := teacher1

	teacher2.Name = "Ki Hajar Dewantoro"

	fmt.Println(teacher1)
	fmt.Println(teacher2)
}

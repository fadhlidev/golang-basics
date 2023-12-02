package main

import "fmt"

type Student struct {
	Name string
}

func main() {
	student1 := Student{"John"}
	student2 := &student1

	student2.Name = "Jonny"

	fmt.Println(student1)
	fmt.Println(*student2)
}

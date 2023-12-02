package main

import "fmt"

type Bird struct {
	Name string
}

func (bird Bird) Tweet() {
	fmt.Println(bird.Name, ": Tweet tweet")
}

func main() {
	b := Bird{"X"}
	b.Tweet()
}

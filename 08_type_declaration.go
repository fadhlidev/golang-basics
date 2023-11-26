package main

import "fmt"

func main() {
	type slug string

	var id slug = "this-is-a-slug"
	fmt.Println(id)
	fmt.Println(slug("this-is-another-slug"))
}

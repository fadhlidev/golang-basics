package main

import "fmt"

type HasName interface {
	GetName() string
}

type Product struct {
	name  string
	Price int64
}

func (product Product) GetName() string {
	return product.name
}

func main() {
	p := Product{"Smartphone", 3000000}
	fmt.Println(p.GetName())
}

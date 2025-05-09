package main

import "fmt"

type Pizza struct {
	ID    int
	name  string
	price float64
}

func main() {
	var pizzas = []Pizza{
		{ID: 1, name: "Toscana", price: 49.5},
		{ID: 2, name: "Marguerita", price: 79.5},
		{ID: 3, name: "Atum com queijo", price: 69.5},
	}
	fmt.Println(pizzas)
}

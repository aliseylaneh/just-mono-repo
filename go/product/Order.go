package product

import "math/rand"

type Order struct {
	id   int
	item string
}

func CreateOrder(item string) *Order {
	return &Order{id: rand.Int(), item: item}
}

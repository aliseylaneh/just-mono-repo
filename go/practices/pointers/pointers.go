package main

import "fmt"

type Product struct {
	name    *string
	address string
}

func typeChecker(arg interface{}) {
	switch arg {
	case arg.(Product):
		fmt.Println("Type is product")
	case arg.(*Product):
		fmt.Println("Type is a pointer product")
	}
}
func makePointer[T any](t T) *T {
	return &t
}
func main() {
	product := Product{name: makePointer("name"), address: "test"}
	typeChecker(&product)
}

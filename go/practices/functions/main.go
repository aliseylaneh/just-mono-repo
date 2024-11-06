package main

import "fmt"

func handleUploadUseCase(id int, name string, useCase func(int, string) (int, string, string)) {
	id, name, typeOfStruct := useCase(id, name)
	fmt.Printf("id: %v, name: %v, Type of Struct: %v\n", id, name, typeOfStruct)
}

type Product struct {
	id    int
	name  string
	color string
}

type Journal struct {
	id     int
	name   string
	number int
}

func journalUseCase(id int, name string) (int, string, string) {
	journal := Journal{id: id, name: name, number: 3321}
	return journal.id, journal.name, "Journal"
}

func productUseCase(id int, name string) (int, string, string) {
	product := Product{id: id, name: name, color: "red"}
	return product.id, product.name, "Product"
}

func main() {
	handleUploadUseCase(10, "Journal of New York", journalUseCase)
	handleUploadUseCase(24, "Neapolitan Pizza", productUseCase)
}

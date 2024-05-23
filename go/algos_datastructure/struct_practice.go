package main

import (
	"fmt"
	"time"
)

var autoIncrementedIdentifier int = 1

type Product struct {
	id               int
	name             string
	shortDescription string
	createdAt        time.Time
}

func (product *Product) viewProductDetail() string {
	var productDetails string = "Product Name: " + product.name + "\nProduct Description: " + product.shortDescription
	return productDetails
}
func (product *Product) updateProduct(name string, shortDescription string) {
	product.name = name
	product.shortDescription = shortDescription
}

func createProduct(name string, shortDescription string) *Product {
	currentTime := time.Now()
	var product Product = Product{id: autoIncrementedIdentifier, name: name, shortDescription: shortDescription, createdAt: currentTime}
	autoIncrementedIdentifier++
	return &product
}

func main() {
	name := "Tesla V2"
	shortDescription := "This car company is Tesla and it's model is version 3"
	var product *Product = createProduct(name, shortDescription)
	fmt.Println(product.viewProductDetail())
	product.updateProduct("Tesla V3", "This car company is Tesla and it's model is version 3")
	fmt.Println(product.viewProductDetail())

}

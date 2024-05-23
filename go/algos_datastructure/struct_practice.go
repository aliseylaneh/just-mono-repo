package main

import (
	"fmt"
	"os"
	"time"
)

var autoIncrementedIdentifier int = 1

type Product struct {
	id               int
	name             string
	shortDescription string
	createdAt        time.Time
}

func (product *Product) createProductFile() {
	file, _ := os.Create(product.name + ".txt")
	content := fmt.Sprintf("ID: %v\nName: %v\nDescription: %v\n", product.id, product.name, product.shortDescription)
	file.WriteString(content)
	file.Close()
}
func (product *Product) viewProductDetail() {
	fmt.Printf("ID:%v,Product Name: %v\nProduct Description: %v\n", product.id, product.name, product.shortDescription)

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

func productStruct() {
	name := "Tesla V2"
	shortDescription := "This car company is Tesla and it's model is version 2"
	var product *Product = createProduct(name, shortDescription)
	product.viewProductDetail()
	product.updateProduct("Tesla V3", "This car company is Tesla and it's model is version 3")
	product.viewProductDetail()
	product.createProductFile()
}

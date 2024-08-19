package main

import (
	"fmt"
	"os"
	"time"
)

var autoIncrementedIdentifier int = 1

type Color struct {
	name string
	code string
}
type Product struct {
	id               int
	name             string
	shortDescription string
	color            *Color
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
	var color Color = Color{name: "Blue", code: "123HXF"}
	var product Product = Product{id: autoIncrementedIdentifier,
		name:             name,
		shortDescription: shortDescription,
		createdAt:        currentTime,
		color:            &color}
	autoIncrementedIdentifier++
	return &product
}

func (product *Product) printProductDetails(anotherProduct *Product) {
	fmt.Printf("This is product number one with number of %v and description of %v", anotherProduct.name, anotherProduct.shortDescription)
	fmt.Printf("This product name is %v and the descritpion is %v\n", product.name, product.shortDescription)
}
func (product *Product) getColor() {
	fmt.Printf("Product color is %v,\nColor code is %v\n", product.color.name, product.color.code)
}

func productStruct() {

	name := "Tesla V2"
	shortDescription := "This car company is Tesla and it's model is version 2"
	var product *Product = createProduct(name, shortDescription)
	product.viewProductDetail()
	product.updateProduct("Tesla V3", "This car company is Tesla and it's model is version 3")
	product.viewProductDetail()
	product.createProductFile()
	var productNumberTwo *Product = createProduct("Tesla V3", "This is another models of Tesla")
	productNumberTwo.viewProductDetail()
	productNumberTwo.printProductDetails(product)
	product.getColor()
}

package product

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
	orders           []Order
}

func (product *Product) createProductFile() {
	file, _ := os.OpenFile(product.name+".txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	content := fmt.Sprintf("ID: %v\nName: %v\nDescription: %v\n", product.id, product.name, product.shortDescription)
	file.WriteString(content)
	file.Close()
}
func (product *Product) ViewProductDetail() {
	fmt.Printf("ID:%v,Product Name: %v\nProduct Description: %v\n", product.id, product.name, product.shortDescription)

}
func (product *Product) updateProduct(name string, shortDescription string) {
	product.name = name
	product.shortDescription = shortDescription
}

func CreateProduct(name string, shortDescription string) *Product {
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

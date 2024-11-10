package main

import "fmt"

type Logger interface {
	Log()
}

type Product struct {
	name string
}

func (product *Product) Log() {
	fmt.Println(product.name)
}

type customType string

func (cType customType) Log() {
	fmt.Println(cType)
}

func printLog(data Logger) {
	data.Log()
}
func main() {
	var myType customType = "Ali Seylaneh"
	printLog(myType)

	product := Product{name: "Car"}
	printLog(&product)
}

package product

import "fmt"

type ProductAggregate struct {
	id      int
	product *Product
}

func NewProductAggregate(product *Product) *ProductAggregate {
	return &ProductAggregate{id: autoIncrementedIdentifier, product: product}
}

func (aggregate *ProductAggregate) ChangeOrders(order []Order) {
	aggregate.product.orders = order
}

func (aggregate *ProductAggregate) GetProductOrders() {
	fmt.Println(aggregate.product.orders)
}
func (aggregate *ProductAggregate) GetProductDetail() {
	fmt.Println(aggregate.id, aggregate.product.id, aggregate.product.name, aggregate.product.orders)
}

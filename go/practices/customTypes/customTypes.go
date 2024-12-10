package main

import (
	"errors"
	"fmt"
)

type Person struct {
	name  string
	phone int
}

type personData map[string]Person
type customInt int

func (number customInt) pow(power int) customInt {
	result := number
	for i := 1; i < power; i++ {
		result *= number
	}
	return result
}

type Inventory struct {
	id    string
	stock uint
}

type inventories map[string]Inventory

type getInventories func(id string) (Inventory, error)

func main() {
	//var people personData = personData{
	//	"Person Number one": Person{name: "Ali", phone: 2311},
	//}
	//fmt.Println(people)
	//var myNumber customInt = 2
	//result := myNumber.pow(2)
	//fmt.Println(result)
	var myFunc getInventories = func(id string) (Inventory, error) {
		listOfInventory := []Inventory{{id: "SKI-TEST", stock: 32}, {id: "SKU-2", stock: 44}}
		for _, inventory := range listOfInventory {
			if inventory.id == id {
				return inventory, nil
			}
		}
		return Inventory{}, errors.New("Inventory not found !!")
	}

	inventory, err := myFunc("SKO-TEST")
	if err == nil {
		fmt.Println(inventory)
	} else {
		fmt.Println(err)
	}

}

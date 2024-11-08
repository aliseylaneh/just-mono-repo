package main

import "fmt"

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
func main() {
	var people personData = personData{
		"Person Number one": Person{name: "Ali", phone: 2311},
	}
	fmt.Println(people)
	var myNumber customInt = 2
	result := myNumber.pow(2)
	fmt.Println(result)

}

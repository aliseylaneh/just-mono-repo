package main

import "fmt"

func TryingNewFunction() {
	justMyValue := new(int)
	fmt.Println(justMyValue)
	fmt.Println(*justMyValue)

	anotherValue := new(map[string]string)
	fmt.Println(anotherValue)
	fmt.Println(*anotherValue)
}

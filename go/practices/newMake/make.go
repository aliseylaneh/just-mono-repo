package main

import (
	"fmt"
)

func main() {

	names := make([]string, 2, 3)
	names[0] = "Jakob"
	names[1] = "Henry"
	names = append(names, "Jake")

	fmt.Println(names)

}

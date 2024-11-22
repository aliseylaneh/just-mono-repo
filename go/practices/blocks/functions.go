package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5 // This will shadow the upper block variable x, but it's not going to be accessible in another blocks
		// Also if you shadow a variable you can't access the original variable at all
		fmt.Println(x)
	}
	fmt.Println(x)
	// Shadowing identifiers from universe block in go is forbidden and can cause various problems in your program
	fmt.Println(true)
	true := 10 // This is forbidden
	fmt.Println(true)
	// if statement
	if n := rand.Intn(10); n == 0 { // You can declare variable depended and accessible on if condition scope
		fmt.Printf("%v is low", n)
	} else if n > 5 {
		fmt.Printf("%v is too big", n)

	} else {
		fmt.Printf("%v is suitable number", n)
	}
	fmt.Println()
	i := 5
	for ; i >= 0; i-- {
		//fmt.Println(i)
	}
	evenVals := []int{2, 4, 6, 8, 10, 12} // evenVals declared and not used
	_ = evenVals
	customers := map[string]string{
		"Ali":  "Developer",
		"Jake": "Product manager",
		"John": "QA",
	}

	for name, position := range customers {
		fmt.Println(name, position)

	}
	samples := []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
}

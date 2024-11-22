package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5 // This will shadow the upper block variable x, but it's not going to be accessible in another blocks
		// Also if you shadow a variable you can't access the original variable at all
		fmt.Println(x)
	}
	fmt.Println(x)
}

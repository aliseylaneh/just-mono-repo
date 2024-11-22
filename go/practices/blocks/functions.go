package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5 // This will shadow the upper block variable x, but it's not going to be accessible in another blocks
		fmt.Println(x)
	}
	fmt.Println(x)
}

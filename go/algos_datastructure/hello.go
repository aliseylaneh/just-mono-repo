package main

import "fmt"

func main() {
	var greetingText string = "Hello Go"
	counter := 10
	for i := 0; i < 5; i++ {
		counter += 10
		fmt.Println(greetingText)
	}
	finalCounter := counter
	print(finalCounter)
}

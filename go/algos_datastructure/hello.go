package main

import "fmt"

func main() {
	var greetingText string = "Hello Go"
	counter := 10
	for i := 0; i < counter; i++ {
		fmt.Println(greetingText)
	}
}

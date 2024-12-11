package main

import "fmt"

type Inner struct {
	name string
}

type Outer struct {
	Inner
	name string
}

func main() {
	o := Outer{
		Inner: Inner{"Ali"},
		name:  "Jack",
	}
	fmt.Println(o.name)
	fmt.Println(o.Inner.name)
}

package main

import "fmt"

type TypeOneInterface interface {
	write()
}

type TypeTwoInterface interface {
	get()
}

type EmbeddedInterface interface {
	TypeOneInterface
	TypeTwoInterface
}

func testEmbeddedInterfaces(value EmbeddedInterface) {
	value.get()
	value.write()
}

type TypeOneStruct struct {
	name string
}
type TypeTwoStruct struct {
	userInput string
	*TypeOneStruct
}

func (s *TypeOneStruct) write() {
	fmt.Println(s.name)
}
func testEmbeddedStruct() {
	typeTwoStruct := TypeTwoStruct{userInput: "'git checkout -b feature/newBranch'",
		TypeOneStruct: &TypeOneStruct{name: "Ali Seylaneh"}}
	typeTwoStruct.write()
}

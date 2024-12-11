package main

import (
	"fmt"
	"strconv"
)

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.insert(val)
	} else if val > it.val {
		it.right = it.right.insert(val)
	}
	return it
}

func (it *IntTree) contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.contains(val)
	case val > it.val:
		return it.right.contains(val)
	default:
		return true
	}
}
func (it *IntTree) printElements(mainNode string, position string) {
	// TODO needs refactor
	rightValue := 0
	leftValue := 0
	if mainNode != "" {
		mainNode = fmt.Sprintf("Current position is %v of main node %v, ", position, mainNode)
	} else {
		mainNode = ""
	}
	if it.right != nil {
		rightValue = it.right.val
	}
	if it.left != nil {
		leftValue = it.left.val
	}
	fmt.Printf("%v%v<-%v->%v\n", mainNode, rightValue, it.val, leftValue)
	mainNode = strconv.Itoa(it.val)
	if it.right != nil {
		it.right.printElements(mainNode, "right")
	}
	if it.left != nil {
		it.left.printElements(mainNode, "left")
	}
}
func main() {
	var it *IntTree
	it = it.insert(5)
	it = it.insert(3)
	it = it.insert(10)
	it = it.insert(2)
	fmt.Println(it.contains(2))  // true
	fmt.Println(it.contains(12)) // false
	it.printElements("", "top")

}

package main

import "fmt"

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
func (it *IntTree) printElements() {
	// TODO needs refactor
	fmt.Printf("Value is %v and children are: ->\n", it.val)
	if it.right != nil {
		fmt.Printf("%v:Right:%v\n", it.val, it.right.val)
		it.right.printElements()
	}
	if it.left != nil {
		fmt.Printf("%v:Left:%v\n", it.val, it.left.val)
		it.left.printElements()

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
	it.printElements()

}

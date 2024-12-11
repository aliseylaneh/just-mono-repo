package main

import "fmt"

const (
	AttackHeal   = iota + 1
	AttackDamage = iota + 10
	JustAnotherValue
)

// Role Best practice for simulating an enumeration in go is defining a custom type
// and the defining const with this type and iota
type Role int

const (
	Admin Role = iota + 1
	User
	Customer
	Reviewer
)

func main() {
	fmt.Println(AttackHeal, AttackDamage, JustAnotherValue)
}

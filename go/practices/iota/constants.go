package main

import "fmt"

const (
	AttackHeal   = iota + 1
	AttackDamage = iota + 10
	JustAnotherValue
)

func main() {
	fmt.Println(AttackHeal, AttackDamage, JustAnotherValue)
}

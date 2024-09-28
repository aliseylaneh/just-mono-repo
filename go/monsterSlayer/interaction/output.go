package interaction

import "fmt"

func PrintGreeting() {
	fmt.Println("Monster Slayer")
	fmt.Println("Starting new game...")
	fmt.Println("Good luck.")
}

func ShowAvailableActions(specialAttackAvailable bool) {
	fmt.Print("Please choose your action:\n" +
		"1) Attack Monster\n" +
		"2) Heal\n")
	if specialAttackAvailable {
		fmt.Print("3) Special Attack\n:")
	}
}

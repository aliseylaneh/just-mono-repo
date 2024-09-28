package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAttackAvailable bool) {
	userInput, err := getPlayerInput()
	if err != nil {
		fmt.Printf("%v", err)
	}
	switch userInput {
	case "1":
		fmt.Println("Attack Monster")
	case "2":
		fmt.Println("Heal")
	case "3":
		if !specialAttackAvailable {
			fmt.Println("Special Attack is not available")
		}
	default:
		fmt.Println("Invalid Option")
	}
}

func getPlayerInput() (string, error) {
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput, nil
}

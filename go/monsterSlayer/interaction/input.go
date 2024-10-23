package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAttackAvailable bool) string {
	userInput, err := getPlayerInput()
	if err != nil {
		fmt.Printf("%v", err)
	}
	switch userInput {
	case "1":
		return "ATTACK"
	case "2":
		return "HEAL"
	case "3":
		if !specialAttackAvailable {
			fmt.Println("Special Attack is not available")
			break
		}
		return "SPECIAL_ATTACK"
	default:
		return "INVALID"
	}
	return ""
}

func getPlayerInput() (string, error) {
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput, nil
}

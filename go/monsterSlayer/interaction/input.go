package interaction

import (
	"bufio"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAttackAvailable bool) {
	
}

func getPlayerInput() (string, error) {
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput, nil
}

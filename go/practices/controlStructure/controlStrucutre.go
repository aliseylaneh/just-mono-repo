package controlStructure

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetUserInput() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your age: ")
	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.Replace(userAgeInput, "\n", "", -1)
	userAge, err := strconv.ParseInt(userAgeInput, 0, 64)
	if err != nil {
		myError := errors.New("invalid input")
		return int(userAge), myError
	}
	return int(userAge), err
}

func ValidateAge(userAge *int) {
	if *userAge >= 18 && *userAge < 100 {
		fmt.Println("Welcome to club")
	} else if *userAge >= 100 {
		fmt.Println("You are dead")
	} else {
		fmt.Println("You are not old enough to join the club")
	}
}

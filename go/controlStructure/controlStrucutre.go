package controlStructure

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetUserInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your age: ")
	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.Replace(userAgeInput, "\n", "", -1)
	userAge, _ := strconv.ParseInt(userAgeInput, 0, 64)
	if userAge >= 18 && userAge < 100 {
		fmt.Println("Welcome to club")
	} else if userAge >= 100 {
		fmt.Println("You are dead")
	} else {
		fmt.Println("You are not old enough to join the club")
	}

}

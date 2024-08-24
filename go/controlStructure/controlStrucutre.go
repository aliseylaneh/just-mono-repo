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
	fmt.Println(userAge)

}

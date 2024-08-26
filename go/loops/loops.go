package loops

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getUserInput() (*int, error) {
	reader := bufio.NewReader(os.Stdin)
	inputValue, _ := reader.ReadString('\n')
	inputValue = strings.Replace(inputValue, "\n", "", -1)
	inputNumber, err := strconv.ParseInt(inputValue, 0, 64)
	outputNumber := int(inputNumber)
	if err != nil {
		myError := errors.New("invalid input")
		return &outputNumber, myError
	}
	return &outputNumber, err
}

func calculateSumUpTo() {
	fmt.Print("Enter first value: ")
	firstValue, _ := getUserInput()
	fmt.Print("Enter second value: ")
	secondValue, _ := getUserInput()
	sum := 0
	for i := *firstValue; i <= *secondValue; i++ {
		fmt.Println(i)
		sum = sum + i
	}
	fmt.Println(sum)

}

func StartProgram() {
	fmt.Print("1) Calculate sum of numbers between two numbers\n" +
		"2) Calculate the factorial of a number\n" +
		"Enter your choice: ")
	choice, err := getUserInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	if *choice == 1 {
		calculateSumUpTo()
	} else {
		fmt.Println("Option is not available")
	}
}

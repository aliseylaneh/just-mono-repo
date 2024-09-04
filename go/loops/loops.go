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
	inputValue, _ := reader.ReadString('\r')
	inputValue = strings.Replace(inputValue, "\r", "", -1)
	inputNumber, err := strconv.ParseInt(inputValue, 0, 64)
	outputNumber := int(inputNumber)
	if err != nil {
		fmt.Println(err)
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
func sumUpManually() {
	fmt.Println("* Sum Up Manually Started *")
	isEnteringNumbers := true
	sum := 0
	for isEnteringNumbers {
		fmt.Print("Enter your number:")
		inputValue, err := getUserInput()
		isEnteringNumbers = err == nil
		sum = *inputValue + sum
	}
	fmt.Println(sum)

}
func calculateListSum() {
	fmt.Print("* List Sum Started *\nHow many numbers do you have?: ")
	countOfNumber, err := getUserInput()
	if err != nil {
		return
	}
	var listOfNumbers []int
	for i := 0; i < *countOfNumber; i++ {
		fmt.Printf("Enter your %v number: ", i)
		number, err := getUserInput()
		if err != nil {
			fmt.Println("Invalid Input !!")
			continue
		}
		listOfNumbers = append(listOfNumbers, *number)
	}
	sum := 0
	for index, value := range listOfNumbers {
		fmt.Printf("Index: %v, Value:%v\n", index, value)
		sum = sum + value
	}
	fmt.Println(sum)

}
func StartProgram() {
	fmt.Print("1) Calculate sum of numbers between two numbers\n" +
		"2) Calculate the factorial of a number\n" +
		"3) Calculate range of numbers within a list\n" +
		"Enter your choice: ")
	choice, err := getUserInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	if *choice == 1 {
		calculateSumUpTo()
	} else if *choice == 2 {
		sumUpManually()
	} else if *choice == 3 {
		calculateListSum()
	} else {
		fmt.Println("No Option available")
	}
}

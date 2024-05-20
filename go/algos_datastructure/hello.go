package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	just_a_package "github.com/aliseylaneh/just-mono-repo/just_a_package"
)

func greetings() {
	var greetingText string = "Hello Go"
	counter := 10
	for i := 0; i < 5; i++ {
		counter += 10
		fmt.Println(greetingText)
	}
	finalCounter := counter
	print(finalCounter)
}

func concatinateTypes(defaultDataBaseUrl string, defaultDatabaseUserAndPort string) string {
	concatinated := defaultDataBaseUrl + defaultDatabaseUserAndPort
	return concatinated
}

func runeBoolByte() {
	var emoji rune = 'g'
	var just_a_byte byte = '3'
	just_a_bool := false
	fmt.Println(emoji, just_a_byte, just_a_bool)
}
func printFormatting() {
	my_value := "WoW"
	another_value := "Wowwww"
	fmt.Printf("Your value is %v", my_value)
	im_sprint_f := fmt.Sprintf("%v %v", my_value, another_value)
	fmt.Println(im_sprint_f)
}
func usingAnotherPackage() {
	fmt.Println(just_a_package.GlobalVariable)
}

var reader = bufio.NewReader(os.Stdin)

func testingBufferIo() string {
	userInput, _ := reader.ReadString('\n')
	return userInput
}
func convertToFloatAndPrint(sentence string) {
	sentence = strings.Replace(sentence, "\n", "", -1)
	float_number, exception := strconv.ParseFloat(sentence, 64)
	if exception != nil {
		fmt.Println(float_number, exception)
		return
	}
	fmt.Println(float_number)
}
func changingValueByPoint(age *string) {
	*age = "70"
	fmt.Println(*age)
}
func understandingPointers() {
	age := "23"
	fmt.Println(age)
	changingValueByPoint(&age)
	var userAge *string
	userAge = &age
	*userAge = "28"
	fmt.Println(age)
	*userAge = "665883"
	fmt.Println(*userAge)
	fmt.Println(age)
}

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	var newUser User = User{firstName: "Ali", lastName: "Seylaneh", birthDate: "1998", createdAt: time.Now()}
	fmt.Println(newUser)
}

package main

import (
	"bufio"
	"fmt"
	"os"

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

func testingBufferIo() {
	userInput, _ := reader.ReadString('\n')
	fmt.Println(userInput)
}
func main() {

}

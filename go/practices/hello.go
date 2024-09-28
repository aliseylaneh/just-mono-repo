package main

import (
	"bufio"
	"fmt"
	justAPackage "mono_repo/justAPackage"
	loops "mono_repo/loops"
	"os"
	"strconv"
	"strings"
	"time"
)

func greetings() {
	var greetingText string = "Hello Go"
	counter := 10
	for i := 0; i < 5; i++ {
		counter += 10
		fmt.Println(greetingText)
	}
	finalCounter := counter
	fmt.Println(finalCounter)
}

func concatenateTypes(defaultDataBaseUrl string, defaultDatabaseUserAndPort string) string {
	concatenated := defaultDataBaseUrl + defaultDatabaseUserAndPort
	return concatenated
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
	fmt.Println(justAPackage.GlobalVariable)
}

var reader = bufio.NewReader(os.Stdin)

func testingBufferIo() string {
	userInput, _ := reader.ReadString('\n')
	return userInput
}
func convertToFloatAndPrint(sentence string) {
	sentence = strings.Replace(sentence, "\n", "", -1)
	floatNumber, exception := strconv.ParseFloat(sentence, 64)
	if exception != nil {
		fmt.Println(floatNumber, exception)
		return
	}
	fmt.Println(floatNumber)
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

func (user *User) printUserDetails(justAnArg string) {
	fmt.Println(justAnArg)
	user.firstName = "Jacob"
	fmt.Printf("My name is %v and my last name is %v, my birthdate is %v\n", user.firstName, user.lastName, user.birthDate)
}
func newUser(firstName string, lastName string, birthDate string) *User {
	user := User{firstName: firstName, lastName: lastName, birthDate: birthDate, createdAt: time.Now()}
	return &user
}

func main() {
	////var product *product_package.Product = product_package.CreateProduct("Launch Box", "This is used for your launch in your office")
	////productAggregate := product_package.NewProductAggregate(product)
	////var orders []product_package.Order
	////orderOne := product_package.CreateOrder("Order 1")
	////orderTwo := product_package.CreateOrder("Order 2")
	////orders = append(orders, *orderOne)
	////orders = append(orders, *orderTwo)
	////productAggregate.ChangeOrders(orders)
	////productAggregate.GetProductDetail()
	////justAPackage.ImplementMaps()
	//userAge, err := controlStructure.GetUserInput()
	//if err == nil {
	//	controlStructure.ValidateAge(&userAge)
	//}
	//fmt.Println(userAge, err)
	loops.StartProgram()
}

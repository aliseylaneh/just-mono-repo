package main

import "fmt"

func concatinate_types(defaultDataBaseUrl string, defaultDatabaseUserAndPort string) string {
	concatinated := defaultDataBaseUrl + defaultDatabaseUserAndPort
	return concatinated
}

func main() {
	// var greetingText string = "Hello Go"
	// counter := 10
	// for i := 0; i < 5; i++ {
	// 	counter += 10
	// 	fmt.Println(greetingText)
	// }
	// finalCounter := counter
	// print(finalCounter)
	defaultDataBaseUrl := "psql:localhost@"
	var defaultDatabaseUserAndPort string = "admin:5329"
	fmt.Println(concatinate_types(defaultDataBaseUrl, defaultDatabaseUserAndPort))
}

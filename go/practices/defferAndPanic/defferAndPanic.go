package main

import (
	"fmt"
	"os"
)

func testingOutDeffer() {
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	//	Or you can do below
	defer file.Close()
}

func main() {

}

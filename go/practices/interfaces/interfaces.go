package main

import "fmt"

type Logger interface {
	Log()
}

type Product struct {
	name string
}

func (product *Product) Log() {
	fmt.Println(product.name)
}

type customType string

func (cType customType) Log() {
	fmt.Println(cType)
}

func printLog(data Logger) {
	data.Log()
}
func checkInterfaceTypeAndExecute(value interface{}) {
	switch val := value.(type) {
	case LogWriter:
		err := val.execute()
		if err != nil {
			panic(err)
		}
	case []int:
		newNumbers := append(val, val...)
		fmt.Println("[Info] Just a Slice: ", newNumbers)
	default:
		fmt.Println("[Info] Nothing was processed.")
	}
}
func main() {
	//var myType customType = "Ali Seylaneh"
	//printLog(myType)
	//
	//product := Product{name: "Car"}
	//printLog(&product)
	log := LogWriter{schema: "car", fileType: ".txt", fileName: "car"}
	checkInterfaceTypeAndExecute(nil)
	checkInterfaceTypeAndExecute([]int{1, 2, 3, 4})
	checkInterfaceTypeAndExecute(log)
}

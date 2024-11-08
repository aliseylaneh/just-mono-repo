package main

import "fmt"

type useCaseFunctionArgs func(int, string) (int, string, string)

func handleUploadUseCase(id int, name string, useCaseFunction useCaseFunctionArgs) {
	id, name, typeOfStruct := useCaseFunction(id, name)
	fmt.Printf("id: %v, name: %v, Type of Struct: %v\n", id, name, typeOfStruct)
}

type Product struct {
	id    int
	name  string
	color string
}

type Journal struct {
	id     int
	name   string
	number int
}

func journalUseCase(id int, name string) (int, string, string) {
	journal := Journal{id: id, name: name, number: 3321}
	return journal.id, journal.name, "Journal"
}

func productUseCase(id int, name string) (int, string, string) {
	product := Product{id: id, name: name, color: "red"}
	return product.id, product.name, "Product"
}
func useCase(structName string) (useCaseFunctionArgs, error) {
	switch structName {
	case "product":
		return productUseCase, nil
	case "journal":
		return journalUseCase, nil
	default:
		return nil, fmt.Errorf("faild to find the usecase")
	}

}
func myFunction(args ...string) {
	fmt.Println(args)
}
func main() {
	myFunction()
	journalFunctionUseCase, err := useCase("journal")
	if err != nil {
		return
	}
	productFunctionUseCase, err := useCase("product")
	if err != nil {
		return
	}
	handleUploadUseCase(10, "Journal of New York", journalFunctionUseCase)
	handleUploadUseCase(24, "Neapolitan Pizza", productFunctionUseCase)
}

package justapackage

import "fmt"

func PracticeArrays() {
	var numbers [5]int = [5]int{5, 2, 3, 4, 4}
	newNumbers := [5]int{5, 51, 23, 13, 2}
	fmt.Println(numbers)
	fmt.Println(newNumbers)
	var names [5]string = [5]string{"String", "Mohammad", "Ali", "Saman"}
	fmt.Println(names[0])
	featureNames := names[1:]
	fmt.Println(featureNames)

}

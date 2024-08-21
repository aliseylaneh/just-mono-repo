package justapackage

import "fmt"

func PracticeArrays() {
	var numbers [5]int = [5]int{5, 2, 3, 4, 4}
	newNumbers := [5]int{5, 51, 23, 13, 2}
	fmt.Println(numbers)
	fmt.Println(newNumbers)
	var names [5]string = [5]string{"String", "Mohammad", "Ali", "Saman"}
	fmt.Println(names)
	featureNames := names[:2]
	fmt.Println(featureNames)
	featureNames = featureNames[:3]
	fmt.Println(featureNames)
	fmt.Println(len(featureNames))
	fmt.Println(cap(featureNames))

}
func SlicingArrays() {
	var names []string
	var newNames []string
	names = append(names, "Ali", "Mohammad", "Reza", "Ahmad")
	newNames = names[2:3]
	fmt.Println(newNames)
	fmt.Println(len(newNames), cap(newNames))
	fmt.Println(newNames[0:2])
}
func DynamicArrays() {
	prices := []float64{1, 2}
	var mainPrices []float64
	value := 0.0
	for i := 0; i <= 10; i++ {
		prices = append(prices, value)
		mainPrices = append(prices, value)
		value += 1
		fmt.Println(value, i)
	}
	fmt.Println(mainPrices, prices)
}

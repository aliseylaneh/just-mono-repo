package main

import "fmt"

type Doubler interface {
	Double()
}
type DoubleInt int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := range d {
		d[i] = d[i] * 2
	}
}

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

func main() {
	//var di DoubleInt = 10
	//var di2 DoubleInt = 10
	//var dis = DoubleIntSlice{1, 2, 3}
	//var dis2 = DoubleIntSlice{1, 2, 3}
	//DoublerCompare(&di, &di2)
	//DoublerCompare(&di, dis)
	//DoublerCompare(dis, dis2)
	var i any
	i = 10
	checker, ok := i.(int)
	print(checker, ok)
}

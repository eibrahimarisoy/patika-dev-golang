package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1 : int x, float64 y type conversion sample

	x := 5
	var y float64
	y = float64(x)

	fmt.Printf("%v, %T\n", y, y)

	// 2 : multiple assing sample x, y = y, x

	a := 5
	b := 10

	fmt.Println(a, b)

	a, b = b, a

	fmt.Println(a, b)
	// 3 : non English variable names

	// 4 : shadowing kavramı? gölgeleme
	k := 5
	if true {
		k := 10
		k++
		fmt.Println(k)
	}
	fmt.Println(k) // k=10 iken 11 k=10 iken 5

	// 5 : 40 as a string

	j := 40
	fmt.Println(j) //=>40 int

	l := string(j)
	fmt.Println(l) // (

	i := strconv.Itoa(j)
	fmt.Println(i) // 40 string

}

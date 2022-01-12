package main

import "fmt"

const i = 14

func main() {
	// 1-) x = x - 10 vs x -=10

	// x = x - 10 assignment statement
	// x -= 10 assignment operation

	// 2-) K = F - 32 / 1.8 + 273   => -40 F kaç K derecedir?
	F := -40

	K := float32(F-32)/1.8 + 273

	fmt.Println(K)

	// 4-) Sabitlerde Shadowing Kavramı çalışır mı?
	// x := 24
	// const x int = 24
	const i = 24
	fmt.Println(i)

	// 5-) const x = 4, const y = 5.4,  x + y?

	const k = 4
	const l = 5.4

	fmt.Println(k, l)
	fmt.Println(k + l)

	// 6-) const x float64 = 6.4 , y := 4 + x, y?

	const x float64 = 6.4
	y := 4 + x
	fmt.Println(y)
}

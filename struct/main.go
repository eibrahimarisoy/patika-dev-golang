package main

import "fmt"

// 1 -) Underlying Type 'int' olacak şekilde kendi veri tipinizi oluşturunuz
// veri tipi ve değerini '10' yazdırınız.

// type myType int

// func main() {

// 	var a myType = 10

// 	fmt.Printf("%T, %v ", a, a)

// }

// 2 -) Başlangıç değeri 10 olan bir X değişkeni oluşturun sonrasında
// bulunduğu adres üzerinden y değişkenini tanımlayıp x değerini 20 yapınız.

// func main() {
// 	// x := 10
// 	// fmt.Println(x)

// 	// y := &x
// 	// fmt.Println(y)
// 	// *y = 20
// 	// fmt.Println(x)

// 	// var x int = 10

// 	// var y *int = &x

// 	// *y = 20

// 	// fmt.Println(x)

// }

// 3 -) Underlying Type struct olan Rectangle type oluşturunuz.
// display, area, circumference metodlarını yazınız.
// type rectangleType struct {
// 	a, b int
// }

// func (r rectangleType) display() {
// 	fmt.Println("Kenar uzunlukları", r.a, r.b)
// }

// func (r rectangleType) area() int {
// 	return r.a * r.b
// }

// func (r rectangleType) circumference() int {
// 	return 2 * (r.a + r.b)
// }

// func main() {

// 	test := rectangleType{5, 4}
// 	test.display()
// 	fmt.Println(test.area())
// 	fmt.Println(test.circumference())

// }

// 4-) family aile bireyleri şeklinde veri tipi oluşturalım, underlying struct. Aile bireylerinin hepsini farklı
// şekilde tanımlayınız. Sonrasında for döngüsünde yazdırınız. field age, name, isMarried field.

type family struct {
	name      string
	age       int
	isMarried bool
}

func showFamily() []family {
	family1 := family{
		name:      "Arin",
		age:       5,
		isMarried: false,
	}

	family2 := family{
		name: "Elis",
		age:  3,
	}

	family3 := family{
		"Gurcan",
		40,
		true,
	}

	var family4 family
	family4.name = "Gamze"
	family4.age = 40
	family4.isMarried = true

	return []family{family1, family2, family3, family4}
}

func main() {
	families := showFamily()

	for _, i := range families {
		fmt.Println(i)
	}
}

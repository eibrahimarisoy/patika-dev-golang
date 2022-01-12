package main

import "fmt"

func main() {
	// 1-) studentName --> John Doe, grade --> 77, isPassed --> true değişkenlerini
	// 3 farklı yöntem ile oluşturup, çıktısını yazdırınız.

	var studentName string = "John Doe"
	var grade int8 = 77
	var isPassed bool = true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)

	var studentName1 = "John Doe"
	var grade1 = 77
	var isPassed1 = true

	fmt.Println(studentName1)
	fmt.Println(grade1)
	fmt.Println(isPassed1)

	studentName2 := "John Doe"
	grade2 := 77
	isPassed2 := true

	fmt.Println(studentName2)
	fmt.Println(grade2)
	fmt.Println(isPassed2)

	// 2-) yukarıda belirtilen değişkenleri tek satır içerisinde tanımlayınız.

	var studentName3, grade3, isPassed3 = "John Doe", 77, true

	fmt.Println(studentName3)
	fmt.Println(grade3)
	fmt.Println(isPassed3)

	studentName4, grade4, isPassed4 := "John Doe", 77, true
	fmt.Println(studentName4)
	fmt.Println(grade4)
	fmt.Println(isPassed4)

	// 3-) "Declaration", "Assign", "Initialization", "Initial Value" kavramlarını açıklayınız. (Terminoloji)

	// var studentName5 string = "John Doe"
	studentName5 := "John Doe1"
	studentName5 = "John Doe2"

	fmt.Println(studentName5)

	// 4-) "Statically Typed" vs "Dynamically Typed" ifadelerini GO ve Python üzerinden gösteriniz.

	// 5-) ":=" vs "=" aradaki farkı gösteriniz, double declaration

	var studentName6 string = "John Doe"
	studentName6 = "John Doe1"
	fmt.Println(studentName6)

	studentName7 := "John Doe"
	studentName7 = "John Doe2"
	fmt.Println(studentName7)
}

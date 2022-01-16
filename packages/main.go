package main

import "fmt"

// 1 -) 5 string elemandan oluşan bir array ve 5 int elemandan oluşan slice oluşturup index numaralarıyla birlikte gösterelim.

// func main() {
// 	// var x = [5]int{0, 1, 2, 3, 4}
// 	// var x = [...]int{0, 1, 2, 3, 4}
// 	// x := [5]int{0, 1, 2, 3, 4}
// 	x := [...]int{0, 1, 2, 3, 4}

// 	fmt.Println(x)

// 	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
// 	fmt.Println(cars)

// 	for i, car := range cars {
// 		fmt.Println(i, car)
// 	}

// 	for i, num := range x {
// 		fmt.Println(i, num)
// 	}

// }

// 2 -) [0,1,2,3,4,5,6,7,8] slice dan [0,1,2,3], [4,5,6], [6,7,8], [2,3,6,7] slicelarını oluşturunuz.

// func main() {

// 	var main_slc = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}

// 	first_slc := main_slc[0:4]

// 	fmt.Println(first_slc)

// 	second_slc := main_slc[4:7]

// 	fmt.Println(second_slc)

// 	third_slc := main_slc[6:]

// 	fmt.Println(third_slc)

// 	forth_slc := append(main_slc[2:4], main_slc[6:8]...)

// 	fmt.Println(forth_slc)

// }

// 3 -) slicelar için copy metodunu ve assign ( = ) ile farkını açıklayınız.

// func main() {
// 	mySlc := []int{1, 2, 3}
// 	mySlc2 := make([]int, 2)
// 	fmt.Println(mySlc)
// 	fmt.Println(mySlc2)

// 	copy(mySlc2, mySlc)
// 	fmt.Println(mySlc)
// 	fmt.Println(mySlc2)

// 	mySlc[0] = 100
// 	fmt.Println(mySlc)
// 	fmt.Println(mySlc2)
// }

// 4 -) map gösterimi ile yazar ve yazarlara ait kitapların isimlerini for range ile gösterin.

// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
// b := map[KeyType]ValueType{key1:value1, key2:value2,...}

// myMap := map[string][]string{
// 	"Yaşar Kemal":     []string{"İnce Memed", "Yer Demir Gök Bakır"},
// 	"Sabahattin Ali":  []string{"Kuyucaklı Yusuf", "Kürk Mantolu Madonna", "Değirmen"},
// 	"Haruki Murakami": []string{"1Q84", "Dans Dans Dans", "Kumandanı Öldürmek"},
// }
func main() {
	myMap := map[string][]string{
		"Yaşar Kemal":     {"İnce Memed", "Yer Demir Gök Bakır"},
		"Sabahattin Ali":  {"Kuyucaklı Yusuf", "Kürk Mantolu Madonna", "Değirmen"},
		"Haruki Murakami": {"1Q84", "Dans Dans Dans", "Kumandanı Öldürmek"},
	}
	for i, key := range myMap {
		fmt.Print(i, " ")
		for _, data := range key {
			fmt.Print(data)
		}
		fmt.Println()
	}

}

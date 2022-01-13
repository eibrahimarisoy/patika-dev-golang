package main

import "fmt"

// 1-) 1 ile 10 arasındaki sayıları if yapısıyla tek - çift olarak yazdırınız.

// func main() {

// 	for i := 1; i < 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i, "çifttir")
// 			continue
// 		}
// 		fmt.Println(i, "tektir")
// 	}

// }

// 2-) for yapısını kullanarak Go'da olmayan while döngüsüne örnek veriniz.

// func main() {
// 	i := 5
// 	for i < 10 {

// 		fmt.Println(i)
// 		i++

// 	}
// }

// 3-) Switch fallthrough ifadesini açıklayınız.

// func main() {
// 	i := 50

// 	switch i {
// 	case 49:
// 		fmt.Println("49")
// 	case 50:
// 		fmt.Println("50")
// 		fallthrough
// 	default:
// 		fmt.Println("default")
// 	}

// }

// 4-) Aşağıdaki if döngüsünü daha idiomatic hale getiriniz.

/* package main
import (
	"fmt"
)
func main() {
	if x := 20; x%2 == 0 {
		fmt.Println(x, "çifttir")
	} else {
		fmt.Println(x, "tektir")
	}
} */

// func main() {
// 	x := 20
// 	if x%2 == 0 {
// 		fmt.Println(x, "çifttir")
// 		return
// 	}
// 	fmt.Println(x, "tektir")
// }

// 5-) 1 ile 50 arasındaki asal sayıları gösteren bir program yazınız.

func main() {

	for i := 2; i < 50; i++ {
		asal := false

		for j := 2; j < i/2; j++ {
			if i%j == 0 {
				asal = true
				break
			}
		}
		if !asal {
			fmt.Println(i)
		}
	}
}

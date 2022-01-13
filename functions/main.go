package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// 1 -) Iki rakam arasında toplama, çıkarma ve çarpma
// işleminin yapıldığı bir fonkiyon yazınız.

// func main() {

// 	x := 50
// 	y := 40

// 	sum, dif, prod := calculate(x, y)

// 	fmt.Println("Toplam: ", sum)
// 	fmt.Println("Fark: ", dif)
// 	fmt.Println("Çarpım: ", prod)

// }

// func calculate(x, y int) (int, int, int) {
// 	sum := x + y
// 	dif := x - y
// 	prod := x * y

// 	return sum, dif, prod
// }

// 2 -) Kullanıcı tarafından girilen nota göre geçtiniz
// veya kaldınız geri dönüşünü yazdırınız.

// func main() {

// 	grade, _ := getGrade()

// 	fmt.Printf("%T\n", grade)
// 	fmt.Println(grade)

// 	if grade > 50 {
// 		fmt.Println("Geçtiniz")
// 		return
// 	} else {
// 		fmt.Println("Kaldınız")
// 	}

// }

// func getGrade() (int, error) {
// 	reader := bufio.NewReader(os.Stdin)
// 	input, err := reader.ReadString('\n')

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	input = strings.TrimSpace(input)
// 	num, err := strconv.Atoi(input)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("%T\n", num)

// 	return num, err
// }

// 3 -) 1 ile yüz arasındaki bir sayıyı tahmin etme uygulaması
// yazınız. Toplam tahmin hakkınız 10 olsun.

func main() {
	rand := numRand(1, 100)
	fmt.Println(rand)

	fmt.Println("1 ile 100 arasındaki sayıyı bulmaya çalışın")
	reader := bufio.NewReader(os.Stdin)

	for attempt := 0; attempt < 10; attempt++ {

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}

		if num > rand {
			fmt.Println("Daha küçük sayı giriniz")
		} else if num < rand {
			fmt.Println("Daha büyük sayı giriniz")
		} else {
			fmt.Println("Bingo")
			break
		}
	}

}

func numRand(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min
}

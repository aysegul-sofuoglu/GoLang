package loops

import "fmt"

func Demo5() {

	sayi1 := 220
	sayi2 := 284
	toplam1 := 0
	toplam2 := 0

	for i := 1; i < sayi1; i++ {

		if sayi1%i == 0 {
			toplam1 = toplam1 + i

		}
	}

	for j := 1; j < sayi2; j++ {

		if sayi2%j == 0 {
			toplam2 = toplam2 + j

		}
	}

	if toplam1 == sayi2 && toplam2 == sayi1 {
		fmt.Println("arkadaş sayılar")
	} else {
		fmt.Println("arkadaş sayı değiller")
	}
}

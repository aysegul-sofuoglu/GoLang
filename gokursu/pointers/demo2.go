package pointers

import "fmt"

func Demo2(sayilar []int) {
	sayilar[0] = 100
	fmt.Println("demodaki sayı: ", sayilar[0])

	//sadece sayısal verilerde bellek alanıyla çalışılır.
	//diziler maplar gibi konularda pointera gerek yok
}

package pointers

import "fmt"

func Demo1(sayi *int) {

	*sayi = *sayi + 1
	fmt.Println("demodaki sayının bellekteki yeri: ", sayi) //* koymazsak değişkenin bellekteki adresini verir
	fmt.Println("demodaki sayı", *sayi)
}

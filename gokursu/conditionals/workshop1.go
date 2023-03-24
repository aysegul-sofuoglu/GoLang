package conditionals

import "fmt"

func Demo3() {

	var sayi1 int = 10
	var sayi2 int = 50
	var sayi3 int = 20

	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2
	}
	if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}

	fmt.Printf("en büyük sayı: %v", enBuyuk)

}

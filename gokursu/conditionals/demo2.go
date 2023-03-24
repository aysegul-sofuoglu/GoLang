package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilecek float64 = 1600

	if cekilecek > hesap {

		fmt.Println("yetersiz bakiye")
	} else {
		hesap = hesap - cekilecek
		fmt.Println("kalan: " + fmt.Sprintf("%v", hesap))
	}
}

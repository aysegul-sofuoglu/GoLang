package conditionals

import "fmt"

func Demo1() {

	var hesap float64 = 1000
	var cekilecek float64 = 600

	if cekilecek > hesap {

		fmt.Println("yetersiz bakiye")
	}
	if cekilecek <= hesap {

		hesap = hesap - cekilecek
		fmt.Println("kalan: " + fmt.Sprintf("%v", hesap))
	}
}

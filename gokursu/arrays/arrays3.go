package arrays

import "fmt"

func Demo3() {

	sayilar := [5]int{5, 15, 10, 9, 21}
	fmt.Println(sayilar)

	enBuyuk := sayilar[0]

	for i := 0; i < len(sayilar); i++ {

		if sayilar[i] > enBuyuk {
			enBuyuk = sayilar[i]
		}

	}

	fmt.Println("en büyük sayı: ", enBuyuk)
}

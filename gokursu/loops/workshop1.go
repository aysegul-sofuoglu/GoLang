package loops

import "fmt"

func Demo3() {

	var sayi int = 55
	var tahmin int = 0
	var sayac int = 1

	fmt.Println("Tahmininiz: ")
	fmt.Scanln(&tahmin)

	for tahmin != sayi {

		if tahmin < sayi {
			fmt.Println("yukarı")
			fmt.Scanln(&tahmin)
			sayac++
		}
		if tahmin > sayi {
			fmt.Println("aşağı")
			fmt.Scanln(&tahmin)
			sayac++
		}
	}

	basari := ""
	if sayac > 0 && sayac <= 3 {
		basari = "süper"
	} else if sayac <= 6 {
		basari = "güzel"
	} else {
		basari = "fena değil"
	}

	if tahmin == sayi {
		fmt.Printf("tebrikler bildiniz! %v tahmin: %v", sayac, basari)

	}

}

package loops

import "fmt"

func Demo4() {

	var sayi int = 0
	fmt.Println("bir sayı giriniz: ")
	fmt.Scanln(&sayi)

	asalMi := true

	for i := 2; i < sayi; i++ {

		if sayi%i == 0 {

			asalMi = false
		}
	}
	if asalMi == true {
		fmt.Println("asaldır.")
	} else {
		fmt.Println("asal değildir.")
	}

}

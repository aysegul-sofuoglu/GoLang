package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklimdakiSayi := 50
	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz.")
	}

	if tahmin > aklimdakiSayi {
		return "Daha küçük bir sayı gir.", nil
	}

	if tahmin < aklimdakiSayi {
		return "Daha büyük bir sayı gir.", nil
	}
	return "Bildiniz", nil
}

func Demo2() {
	mesaj, hata := TahminEt(55)
	fmt.Println(mesaj)
	fmt.Println(hata)
	// fmt.Println(TahminEt(120))
	// fmt.Println(TahminEt(-12))
}

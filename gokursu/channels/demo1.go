package channels

import (
	"fmt"
	"time"
)

//fonksiyonların aynı zamanda çalışması için channel kullanılır

func CiftSayilar(ciftSayiCn chan int) {

	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("çift toplama çalışıyor")
		time.Sleep(1 * time.Second)
	}
	ciftSayiCn <- toplam
}

func TekSayilar(tekSyiCn chan int) {

	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("tek toplama çalışıyor")
		time.Sleep(1 * time.Second)
	}
	tekSyiCn <- toplam
}

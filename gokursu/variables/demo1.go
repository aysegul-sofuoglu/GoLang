package variables

import "fmt"

func Demo1() {

	fmt.Println("hello go!")

	sayi := 2.5
	fmt.Printf("tip: %T\n", sayi)
	var durum bool

	var isim1 = "Ayşe"
	var isim2 = "Ahmet"

	durum = isim1 == isim2
	fmt.Println(durum)
}

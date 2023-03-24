package defer_statement

import "fmt"

func A() {
	fmt.Println("a çalıştı")
}

func C() {
	fmt.Println("c çalıştı")
}

func D() {
	fmt.Println("d çalıştı")
}

func B() {

	defer A() // defer fonksiyon nerde olursa olsun en son çalışır
	defer C() //ilk giren son çıkar mantığı var  (sıra: b-d-c-a)
	defer D()
	fmt.Println("b çalıştı") //önce b çalışır

}

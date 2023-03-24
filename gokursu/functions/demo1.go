package functions

import "fmt"

func Selam(isim string) {
	fmt.Println("merhaba ", isim)
}

func Topla(say1 int, say2 int) (int, string) {

	var toplam = say1 + say2
	bise := "asd"
	return toplam, bise
}

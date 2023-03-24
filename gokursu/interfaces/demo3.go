package interfaces

import "fmt"

func Dogrula(i interface{}) {
	sayi, ok := i.(int) //tip doğrulama
	fmt.Println(sayi, ok)
	//ok true veya false döner

}

func Demo3() {

	//1.
	Dogrula(5)

	//2.
	var say1 interface{} = 5
	Dogrula(say1)

	//iki seçenek aynı

	var say2 interface{} = "ayşegül" //int yerine string girdik hata alıcaz
	Dogrula(say2)

	var say3 interface{} = 5
	Dogrula(say3)

}

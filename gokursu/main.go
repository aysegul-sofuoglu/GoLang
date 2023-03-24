package main

import (
	"fmt"
	"golesson/project"
)

func main() {

	//variables.Demo1()
	//fmt.Println("---------")

	//................................................

	//conditionals.Demo2()
	//fmt.Println("---------")

	//................................................

	//conditionals.Demo3()
	//fmt.Println("---------")

	//................................................

	//loops.Demo1()
	//fmt.Println("---------")

	//................................................

	//loops.Demo2()
	//fmt.Println("---------")

	//................................................

	//loops.Demo3()
	//fmt.Println("---------")

	//................................................

	//loops.Demo4()
	//fmt.Println("---------")

	//................................................

	//loops.Demo5()
	//fmt.Println("---------")

	//................................................

	//arrays.Demo2()
	//fmt.Println("---------")

	//................................................

	//arrays.Demo3()
	//arrays.Demo4()

	//................................................

	//slices.Demo1()
	//slices.Demo2()

	//................................................

	//functions.Selam("ayşe")

	//var sonuç = functions.Topla(4, 7)
	//fmt.Println(sonuç)

	// toplam, cikarma, carpim, bolum := functions.DortIslem(9, 3)
	// fmt.Println("toplam: ", toplam)
	// fmt.Println("çıkarma: ", cikarma)
	// fmt.Println("çarpım", carpim)
	// fmt.Println("bölüm: ", bolum)
	// fmt.Println("---------------------")
	// toplama, _, carpma, _ := functions.DortIslem(9, 3) //alt çizgi koyarsak o sonuç verilmez
	// fmt.Println("toplam: ", toplama)
	// //fmt.Println("çıkarma: ", cikarma)
	// fmt.Println("çarpım", carpma)
	// //fmt.Println("bölüm: ", bolum)

	// fmt.Println(functions.ToplaVariadic(1, 7, 9, 5))
	// fmt.Println(functions.ToplaVariadic(1, 5))

	// dizi := []int{4, 9, 7, 3}
	// fmt.Println(functions.ToplaVariadic(dizi...))

	//................................................

	//maps.Demo1()

	//................................................

	//for_range.Demo1()
	//for_range.Demo2()
	//for_range.Demo3()

	//................................................

	// sayi := 20
	// pointers.Demo1(&sayi)
	// fmt.Println("maindeki sayı: ", sayi)

	// sayilar := []int{1, 2, 3}
	// pointers.Demo2(sayilar)
	// fmt.Println("maindeki sayı: ", sayilar[0])

	//................................................

	//structs.Demo1()
	//structs.Demo2()

	//................................................

	// go goroutines.CiftSayilar() //fonksiyonların sırayla değil eş zamanlı çalışması için
	// go goroutines.TekSayilar()
	// time.Sleep(5 * time.Second)
	// fmt.Println("bitti")

	//................................................

	// ciftToplamCn := make(chan int)
	// tekToplamCn := make(chan int)
	// go channels.CiftSayilar(ciftToplamCn)
	// go channels.TekSayilar(tekToplamCn)

	// ciftSayiToplam, tekSayiToplam := <-ciftToplamCn, <-tekToplamCn

	// carpim := ciftSayiToplam * tekSayiToplam
	// fmt.Println("çarpım: ", carpim)

	//................................................

	//interfaces.Demo1()
	//interfaces.Demo2()
	//interfaces.Demo3()

	//................................................

	//defer_statement.B()
	//defer_statement.Demo2()

	//................................................

	//error_handling.Demo1()
	//error_handling.Demo2()
	//fmt.Println(error_handling.TahminEt2(55))

	//string_functions.Demo1()
	//string_functions.Demo2()

	//restful.Demo1()
	//restful.Demo2()

	products, _ := project.GetAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i])
	}

	// product, _ := project.AddProduct()
	// fmt.Println(product)
}

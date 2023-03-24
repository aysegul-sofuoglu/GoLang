package loops

import "fmt"

func Demo1() {

	var metin string = "Hello"

	i := 1 //i değişkeni tanımladım

	for i <= 5 {

		fmt.Println(metin)
		i = i + 1
	}
	fmt.Println("bitti")
}

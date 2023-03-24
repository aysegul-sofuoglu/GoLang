package slices

import "fmt"

func Demo1() {

	isimler := make([]string, 3)

	isimler[0] = "ayşe"
	isimler[1] = "buse"
	isimler[2] = "ceyda"

	fmt.Println(isimler)

	isimler = append(isimler, "çağla")
	fmt.Println(isimler)
}

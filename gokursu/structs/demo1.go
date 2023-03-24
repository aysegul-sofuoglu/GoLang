package structs

import "fmt"

func Demo1() {

	fmt.Println(product{"bilgisayar", 500, "asus"})
	fmt.Println(product{"bilgisayar", 500, "apple"})
	fmt.Println(product{name: "bilgisayar", unitPrice: 500})
}

type product struct { //class gibi
	name      string
	unitPrice float64
	brand     string
}

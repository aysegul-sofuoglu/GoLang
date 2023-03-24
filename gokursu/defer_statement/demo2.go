package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) Save() {
	fmt.Println("ürün kaydedildi: ", p.productName)
	defer Log() //defer en çok log işlemlerinde kullanılır
}

func Log() {
	fmt.Println("loglandı")
}
func Demo2() {
	p := product{productName: "laptop", unitPrice: 500}
	defer p.Save()
	fmt.Println("bitti")
}

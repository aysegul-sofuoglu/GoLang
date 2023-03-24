package string_functions

import (
	"fmt"
	s "strings" //alias
)

func Demo1() {
	isim := "mustik"
	fmt.Println(s.Count(isim, "y"))    //büyük küçük harf duyarlı
	fmt.Println(s.Contains(isim, "b")) //var mı yok mu
	fmt.Println(s.Index(isim, "l"))    //aranan harfin indexi

	sonuc := s.Index(isim, "l")

	if sonuc != -1 {
		fmt.Println("Harf mevcut")
	} else {
		fmt.Println("Harf mevcut değil")
	}

	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))

}

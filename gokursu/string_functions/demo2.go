package string_functions

import (
	"fmt"
	s "strings" //alias
)

func Demo2() {

	isim := "aysegul"
	fmt.Println(s.HasPrefix(isim, "ays"))
	fmt.Println(s.HasSuffix(isim, "gul"))
	fmt.Println(s.Index(isim, "gul"))

	harfler := []string{"a", "y", "s", "e"}
	//fmt.Println((s.Join(harfler, ""))) //elemanları birleştirir

	sonuc := s.Join(harfler, "-")
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "-", "+", -1)) //-1dersek hepsi değişir
	fmt.Println(s.Replace(sonuc, "-", "+", 2))  //2 tanesi değişir

	kelimeler := []string{"elma", "muz", "çilek"}
	strkelime := s.Join(kelimeler, ",")
	fmt.Println(s.Split(strkelime, ",")) //kelimeleri virgüllere göre ayır

	fmt.Println(s.Repeat(isim, 3))
}

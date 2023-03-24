package slices

import "fmt"

func Demo2() {

	sehirler := []string{"istanbul", "ankara", "izmir"}

	fmt.Println(sehirler)

	kopya := make([]string, len(sehirler))
	fmt.Println(kopya)

	copy(kopya, sehirler)
	fmt.Println(kopya)

	kopya = append(kopya, "adana")
	kopya = append(kopya, "mersin")
	fmt.Println(kopya)

	fmt.Println(kopya[1:3]) //1den 3e kadarki elemanları yaz
	fmt.Println(kopya[1:])  //1den sona kadar
	fmt.Println(kopya[:3])  //en baştan 3e kadar
}

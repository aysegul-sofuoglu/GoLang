package arrays

import "fmt"

func Demo2() {

	var sehirler [5]string
	sehirler[0] = "adana"
	sehirler[1] = "mersin"
	sehirler[2] = "istanbul"
	sehirler[3] = "ankara"
	sehirler[4] = "izmir"

	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}
}

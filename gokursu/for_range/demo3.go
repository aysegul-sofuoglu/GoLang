package for_range

import "fmt"

func Demo3() {
	sozluk := map[string]string{"glass": "bardak", "microphone": "mikrofon"}

	for k, v := range sozluk {
		fmt.Print(k + " : ")
		fmt.Println(v)
	}
}

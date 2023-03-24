package maps

import "fmt"

func Demo1() {

	sozluk := make(map[string]string)
	sozluk["table"] = "masa"
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "kalem"

	fmt.Println(sozluk["book"])
	fmt.Println(sozluk)
	delete(sozluk, "book")
	fmt.Println(sozluk)

	deger, varMi := sozluk["asd"]
	fmt.Println(deger)
	fmt.Println("listede olma durumu: ", varMi)

	sozluk2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(sozluk2)
}

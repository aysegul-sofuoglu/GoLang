package functions

func DortIslem(say1 int, say2 int) (int, int, int, float32) {

	toplama := say1 + say2
	cikarma := say1 - say2
	carpma := say1 * say2
	bolme := float32(say1 / say2)

	return toplama, cikarma, carpma, bolme
}

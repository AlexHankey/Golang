package iteration

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 10; i++ {
		repeated += character
	}
	return repeated
}

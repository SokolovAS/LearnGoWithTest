package iteration

func Repeat(character string, count int) string {
	var repided string

	for i := 0; i < count; i++ {
		repided += character
	}

	return repided
}

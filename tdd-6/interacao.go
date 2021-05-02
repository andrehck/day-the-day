package interacao

func Repetir(a string) string {
	var repeticoes string

	for i := 0; i < 5; i++ {
		repeticoes = repeticoes + a
	}
	return repeticoes
}

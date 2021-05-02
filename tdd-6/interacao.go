package interacao

func Repetir(a string, qtdRpt int) string {
	var repeticoes string

	for i := 0; i < qtdRpt; i++ {
		repeticoes += a
	}
	return repeticoes
}

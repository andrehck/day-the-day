package interacao

const qtdRpt = 5

func Repetir(a string) string {
	var repeticoes string

	for i := 0; i < qtdRpt; i++ {
		repeticoes += a
	}
	return repeticoes
}

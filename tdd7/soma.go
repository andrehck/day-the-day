package main

func Soma(numeros []int) int {
	soma := 0
	for _, numeros := range numeros { //o range permite que percorremos um array,sempre que chamado retorna o indice e o valor.
		soma += numeros
	}
	return soma
}
func SomaTudo(numerosSoma ...[]int) []int {
	var somas []int
	for _, numeros := range numerosSoma {
		somas = append(somas, Soma(numeros))
	}
	return somas
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			final := numeros[1:]
			somas = append(somas, Soma(final))
		}
	}

	return somas
}

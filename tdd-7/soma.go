package main

func Soma(numeros []int) int {
	soma := 0
	for _, numeros := range numeros { //o range permite que percorremos um array,sempre que chamado retorna o indice e o valor.
		soma += numeros
	}
	return soma
}

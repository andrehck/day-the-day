package main

import "fmt"

const ola = "Olá, "
const olaespanhol = "Hola, "
const olafrances = "Bonjour, "
const espanhol = "Espanhol"
const frances = "Francês"

func main() {
	fmt.Println(Ola("mundo", ""))
}

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	return prefixodeSaudacao(idioma) + nome
}

func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = olafrances
	case espanhol:
		prefixo = olaespanhol
	default:
		prefixo = ola
	}
	return

}

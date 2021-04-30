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

	if idioma == espanhol {

		return olaespanhol + nome
	}

	if idioma == frances {
		return olafrances + nome
	}
	return ola + nome

}

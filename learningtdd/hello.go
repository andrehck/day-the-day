package main

import "fmt"

const ola = "Olá, "

func main() {
	fmt.Println(Ola("mundo", ""))
}

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	if idioma == "Espanhol" {
		return "Hola, " + nome
	}
	return ola + nome

}

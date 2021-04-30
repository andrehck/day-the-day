package main

import "fmt"

const ola = "Olá, "

func main() {
	fmt.Println(Ola("mundo"))
}

func Ola(nome string) string {
	if nome == "" {
		nome = "mundo"
	}
	return ola + nome

}

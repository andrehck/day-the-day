package main

import "fmt"

const ola = "Olá, "

func main() {
	fmt.Println(Ola("mundo"))
}

func Ola(nome string) string {
	return ola + nome
}

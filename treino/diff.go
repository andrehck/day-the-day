package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	readfile()
}

func readfile() {
	file, err := ioutil.ReadFile("teste.txt")
	file2, err := ioutil.ReadFile("teste2.txt")
	/*data := make([]byte, 100)
	/*count, err := file.Read(data)*/
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bytes.Compare(file, file2))
}

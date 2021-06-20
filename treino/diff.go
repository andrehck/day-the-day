package main

import (
	"bytes"
	"io/ioutil"
	"log"
)

func diff() int {
	file, err := ioutil.ReadFile("teste.txt")
	file2, err := ioutil.ReadFile("teste2.txt")
	/*data := make([]byte, 100)
	/*count, err := file.Read(data)*/
	if err != nil {
		log.Fatal(err)
	}
	a := (bytes.Compare(file, file2))

	return a
}

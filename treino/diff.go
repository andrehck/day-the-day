package main

import (
	"bytes"
	"io/ioutil"
	"log"
)

func diff() bool {
	file, err := ioutil.ReadFile("teste2.txt")
	file2, err := ioutil.ReadFile("teste.txt")
	/*data := make([]byte, 100)
	/*count, err := file.Read(data)*/
	if err != nil {
		log.Fatal(err)
	}
	a := (bytes.Equal(file, file2))

	return a
}

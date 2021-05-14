package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	a := os.Args[1]
	c := os.Args[2]
	b := runtime.GOOS

	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(b)

	if os.Args[1] == "1" {
		fmt.Print("Ok")
	} else {
		fmt.Print("Fail")
	}
}

func help() {
}

func serviceid() {

}

func address() {

}

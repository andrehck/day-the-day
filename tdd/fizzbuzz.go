package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(i int) string {
	var result string
	var istr = strconv.Itoa(i)

	if i%3 == 0 {
		result = "Fizz"
	}

	if i%5 == 0 {
		result += "Buzz"
	}

	if len(result) == 0 {
		result = istr
	}

	return result
}

func main() {
	fmt.Print("Enter an integer number:")
	var input int
	fmt.Scan("%d, &input")

	for i := 1; i <= input; i++ {
		fmt.Println((fizzbuzz(i)))
	}
}

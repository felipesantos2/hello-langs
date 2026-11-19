package main

import (
	"fmt"

	"examples/hello/morestrings"

	"github.com/google/go-cmp/cmp"
)

// go version
// go env
// go run <file>.go
// go install examples/hello

// entrypoint function like C
func main() {
	fmt.Println("Olá, Mundo!")

	// primitive data types
	// strings
	// boolean
	// numbers
	var myString string
	myString = "Olá, Mundo!"
	fmt.Println(myString)

	var myInt int
	fmt.Println(myInt)

	var myInt8 int8 = 56
	var myInt16 int16 = 10

	fmt.Println(myInt)
	fmt.Println(myInt8)
	fmt.Println(myInt16)

	fmt.Println(morestrings.ReverseRunes("Ovo"))
	fmt.Println(morestrings.ReverseRunes("Banana"))
	fmt.Println(cmp.Diff("Olá M", "Olá, Mundo!"))

}

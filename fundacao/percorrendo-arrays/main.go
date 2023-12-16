package main

import (
	"fmt"
)

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Joshua"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30
	fmt.Println(len(meuArray))
	fmt.Println(meuArray[0])

	for i, v := range meuArray {
		fmt.Printf("O valor de i é %d e o valor de v é %d\n", i, v)
	}
}
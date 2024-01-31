package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(34, 29))

	valor, err := sum(98, 2)
	fmt.Println(valor, err)
}

func sum(a , b int) (int, error) {
	if (a + b > 50) {
		return 0, errors.New("Sum is greater than 50")
	}
	return a + b, nil
}
package main

import "fmt"

func main() {
	salarios := map[string]float64{
		"João": 2000.0,
		"Maria": 3000.0,
	}
	for nome, salario := range salarios {
		fmt.Printf("Nome: %s, Salario: %.2f\n", nome, salario)
	}

	paises := make(map[string]string)
	paises["BR"] = "Brasil"
	fmt.Println(paises)
}
package main

const a = "Hello, World!"

var (
	b bool    = true
	c int     = 10
	d string  = "Joshua"
	e float64 = 1.2
)

func main() {

	//Toda variável declarada localmente deve ser usada
	var a string = "X"
	b := "Y"
	b = "Z"
	println(a)
	println(b)
}
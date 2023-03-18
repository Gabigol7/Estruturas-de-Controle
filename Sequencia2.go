package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Qual valor do numero ? ")
	fmt.Scan(&numero)

	fmt.Println(numero * 2)
	fmt.Println(numero * 3)
	fmt.Println(numero * 4)

}

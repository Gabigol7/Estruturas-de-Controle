package main

import "fmt"

func main() {

	var numero1, numero2, numero3 int

	fmt.Print("Qual o valor do numero1 ? ")
	fmt.Scan(&numero1)

	fmt.Print("Qual valor do numero2 ? ")
	fmt.Scan(&numero2)

	fmt.Print("Qual valor do numero3 ? ")
	fmt.Scan(&numero3)

	fmt.Print(numero1 + numero2 + numero3)

}

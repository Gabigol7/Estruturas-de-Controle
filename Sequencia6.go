package main

import "fmt"

func main() {

	var salário float64

	fmt.Print("Digite seu salário: ")
	fmt.Scan(&salário)

	novoSalario := salário * 1.15

	fmt.Print("salario com aumento é: ", novoSalario)
}

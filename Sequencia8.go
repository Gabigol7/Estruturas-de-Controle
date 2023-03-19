package main

import "fmt"

func main() {

	var diastrabalhados, diaria int

	fmt.Print("diastrabalhados: ")
	fmt.Scan(&diastrabalhados)

	fmt.Print("Digite valor diaria: ")
	fmt.Scan(&diaria)

	fmt.Println("seu salario mensal é: ", diaria*diastrabalhados)
}

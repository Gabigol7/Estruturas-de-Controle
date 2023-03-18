package main

import "fmt"

func main() {

	var peso, altura float64

	fmt.Print("Qual valor do peso ? ")
	fmt.Scanln(&peso)

	fmt.Print("Qual valor da altura ? ")
	fmt.Scanln(&altura)

	//imc := peso / (altura * altura)

	fmt.Println("Seu valor do seu IMC é ", peso/(altura*altura)

}

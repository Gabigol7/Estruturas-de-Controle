package main

import "fmt"

func main() {

	var numero1, numero2, numero3 float64

	fmt.Print("Qual valor do numero1 ? ")
	fmt.Scan(&numero1)

	fmt.Print("Qual valor do numero2 ? ")
	fmt.Scan(&numero2)

	fmt.Print("Qual valor do numero3 ? ")
	fmt.Scan(&numero3)

	//MediaPonderada := (numero1*2 + numero2*3 + numero3*5) / 10

	fmt.Println("A média ponderada dos números é: ", ((numero1*2 + numero2*3 + numero3*5) / 10))

}

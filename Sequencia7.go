package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Digite numero: ")
	fmt.Scan(&numero)

	sucessor := numero + 1

	fmt.Println("O sucessor de ", numero, "é", sucessor)
}

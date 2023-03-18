package main

import "fmt"

func main() {

	var IdadeAnos int

	fmt.Print("Digite a Idade em anos: ")
	fmt.Scan(&IdadeAnos)

	idadeDias := IdadeAnos * 365

	fmt.Print("A idade em dias é: ", idadeDias)

}

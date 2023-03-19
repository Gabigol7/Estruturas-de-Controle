package main

import "fmt"

func main() {

	var produto float64

	fmt.Print("Digite o valor do seu Produto: ")
	fmt.Scan(&produto)

	Produtocomdesconto := produto - (produto * 0.1)

	fmt.Print("Seu produto com desconto é ", Produtocomdesconto)

}

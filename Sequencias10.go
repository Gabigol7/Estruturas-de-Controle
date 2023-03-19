package main

import "fmt"

func main() {

	var libras, quilos float64

	fmt.Print("Digite seu peso em quilos? ")
	fmt.Scan(&quilos)

	libras = quilos * 2.2

	fmt.Print("Seu peso em libras é ", libras)
}

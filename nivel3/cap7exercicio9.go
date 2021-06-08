//- Crie um programa que utilize a declaração switch,
//onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".

package main

import "fmt"

func main() {
	esporteFavorito := "rugby"
	switch esporteFavorito {
	case "rugby":
		fmt.Println("Sim! Meu esporte favorito é Rugby :)")
	default:
		fmt.Println("Não! Esse não é o meu esporte favorito :().")
	}
}

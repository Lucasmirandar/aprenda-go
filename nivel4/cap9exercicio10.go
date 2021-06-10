/*- Utilizando o exercício anterior,
remova uma entrada do map e
demonstre o map inteiro utilizando range.
*/

package main

import "fmt"

func main() {

	x := map[string][]string{
		"Miranda_Lucas": {
			"Jogar", "Treinar",
		},
	}
	x["De Novo_Eu"] = []string{"Estudar", "Dormir"}

	delete(x, "Miranda_Lucas")

	for i, y := range x {
		fmt.Println(i)
		for h, k := range y {
			fmt.Println("\t", h, " ", k)
		}
	}
}

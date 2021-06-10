/*- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.*/

package main

import "fmt"

func main() {

	x := map[string][]string{
		"Miranda_Lucas": {
			"Jogar", "Treinar",
		},
	}
	for i, y := range x {
		fmt.Println(i)
		for h, k := range y {
			fmt.Println("\t", h, " ", k)
		}
	}
}

/*- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores,
utilizando range na slice que contem os sabores de sorvete.*/

package main

import "fmt"

func main() {
	type pessoa struct {
		nome            string
		sobrenome       string
		sorveteFavorito []string
	}

	pessoa1 := pessoa{
		nome:            "Lucas",
		sobrenome:       "Miranda",
		sorveteFavorito: []string{"Chocolate"},
	}
	pessoa2 := pessoa{
		nome:            "Super",
		sobrenome:       "Man",
		sorveteFavorito: []string{"Morango"},
	}
	fmt.Println(pessoa1.nome, pessoa1.sobrenome)

	for _, v := range pessoa1.sorveteFavorito {
		fmt.Println(v)
	}
	fmt.Println(pessoa2.nome, pessoa2.sobrenome)

	for _, v := range pessoa2.sorveteFavorito {
		fmt.Println(v)
	}
}

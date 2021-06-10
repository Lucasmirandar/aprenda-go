/*- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.*/

package main

import "fmt"

func main() {

	type pessoa struct {
		nome            string
		sobrenome       string
		sorveteFavorito []string
	}

	pessoas := make(map[string]pessoa)

	pessoas["Miranda"] = pessoa{
		nome:            "Lucas",
		sobrenome:       "Miranda",
		sorveteFavorito: []string{"Chocolate"},
	}
	pessoas["Man"] = pessoa{
		nome:            "Super",
		sobrenome:       "Man",
		sorveteFavorito: []string{"Morango"},
	}

	for _, v := range pessoas {
		fmt.Println(v.nome, v.sobrenome)
		for _, v := range v.sorveteFavorito {
			fmt.Println("\t", v)
		}
	}

}

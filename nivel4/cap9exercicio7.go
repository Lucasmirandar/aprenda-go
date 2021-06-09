/*- Crie uma slice contendo slices de strings ([][]string). Atribua valores a
este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.*/

package main

import "fmt"

func main() {
	pessoas := [][]string{
		[]string{"Primeira Pessoa:", "Nome,", "Sobrenome,", "Hobby"},
		[]string{"Segunda Pessoa:", "Nome Dois,", "Sobrenome Dois,", "Hobby Dois"},
		[]string{"Terceira Pessoa:", "Nome Três,", "Sobrenome Três,", "Hobby Três"},
	}

	for _, v := range pessoas {
		for _, x := range v {
			fmt.Printf("%v\n", x[:])
		}
	}
}

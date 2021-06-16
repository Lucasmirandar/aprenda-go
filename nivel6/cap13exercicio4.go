/*- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {

	chamada1 := pessoa{
		nome:      "Lucas",
		sobrenome: "Miranda",
		idade:     23,
	}

	chamada1.demonstrar()

}

func (p pessoa) demonstrar() {
	fmt.Println("Nome completo e idade: ", p.nome, p.sobrenome, "-", p.idade, "anos")
}

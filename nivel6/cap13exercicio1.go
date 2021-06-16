/*- Exercício:
  - Crie uma função que retorne um int
  - Crie outra função que retorne um int e uma string
  - Chame as duas funções
  - Demonstre seus resultados*/

package main

import "fmt"

func main() {
	x := retornaInt()
	X, s := retornaIntString()
	fmt.Println(x)
	fmt.Println(X, s)
}

func retornaInt() int {
	return 10
}

func retornaIntString() (int, string) {
	return 10, "Ola"
}

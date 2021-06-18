/*- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.*/

package main

import "fmt"

func retorna() func() {
	return func() {
		fmt.Println("OLA")
	}
}
func main() {
	variavel := retorna()
	variavel()
}

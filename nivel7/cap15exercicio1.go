/*- Crie um valor e atribua-o a uma variável.
- Demonstre o endereço deste valor na memória.*/

package main

import "fmt"

func main() {
	x := 55
	fmt.Println(&x)
}
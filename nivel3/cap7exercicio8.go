// - Crie um programa que utilize a declaração switch, sem switch statement especificado.

package main

import "fmt"

func main() {
	x := 10
	switch {
	case x < 10:
		fmt.Println("X é menor que 10")
	case x == 10:
		fmt.Println("X é igual a 10")
	case x > 10:
		fmt.Println("X é maior que 10")
	default:
		fmt.Println("Erro")
	}
}

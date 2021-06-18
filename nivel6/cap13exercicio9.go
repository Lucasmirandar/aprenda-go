//- Callback: passe uma função como argumento a outra função.

package main

import "fmt"

func main() {
	funcao1(funcao2)
}

func funcao1(f func())  {
	fmt.Println("OLA1")
	f()
}

func funcao2(){
	fmt.Println("OLA2")
}
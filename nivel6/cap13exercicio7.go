/*- Atribua uma função a uma variável.
- Chame essa função.*/

package main

import "fmt"

func main(){
	teste := func(){
		fmt.Println("OLA")
	}
teste()
}
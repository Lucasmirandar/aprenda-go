/*- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
1. 42
2. "James Bond"
3. true
- Agora demonstre os valores nestas variáveis, com:
1. Uma única declaração print
2. Múltiplas declarações print*/

package main

import "fmt"

func main() {
	x, y, z := 42, "James Bond", true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

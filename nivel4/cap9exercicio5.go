/* Comece com essa slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Utilizando slicing e append, crie uma slice y que contenha os valores:
    - [42, 43, 44, 48, 49, 50, 51]*/

package main

import "fmt"

func main() {
	//          0   1   2   3   4   5   6   7  8   9
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := []int{}

	y = append(y, x[0:3]...)

	y = append(y, x[6:10]...)

	fmt.Printf("%s%v\n%s%v", "slice x:", x, "slice y:", y)
}

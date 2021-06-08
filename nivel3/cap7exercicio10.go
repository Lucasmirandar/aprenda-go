/*esultado das expressões:
  - fmt.Println(true && true)
  - fmt.Println(true && false)
  - fmt.Println(true || true)
  - fmt.Println(true || false)
  - fmt.Println(!true)*/

package main

import "fmt"

func main() {
	//true
	fmt.Println(true && true)
	//false
	fmt.Println(true && false)
	//true
	fmt.Println(true || true)
	//true
	fmt.Println(true || false)
	//false
	fmt.Println(!true)
}

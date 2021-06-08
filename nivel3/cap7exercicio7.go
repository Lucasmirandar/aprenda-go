package main

import "fmt"

func main() {
	x := 10
	if x < 10 {
		fmt.Println("X < 10")
	} else if x == 10 {
		fmt.Println("X = 10")
	} else {
		fmt.Println("X > 10")
	}
}

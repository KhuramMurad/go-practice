package main

import (
	"fmt"
)

func main() {
	var a int = 12
	var b int = 12
	fmt.Println(a <= b) // True a = 12. b = 12

	a = 20
	fmt.Println(a <= b) // False a = 20 , b = 12

	b = 100
	fmt.Println(a <= b) // True a = 20 , b = 100

	c := 0
	fmt.Println(a <= c) // False a = 20 , c = 0

}

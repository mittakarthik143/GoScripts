package main

import (
	"fmt"
)

// Miscellaneous operators

func main() {
	var a string = "Go Lang"
	var p *string = &a

	fmt.Println("A is: ", a)
	fmt.Println("A address is: ", &a)
	fmt.Println("P is: ", p)
	fmt.Println("Pointer of P data is: ", *p)
}

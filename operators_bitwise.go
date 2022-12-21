package main

// Bitwiser operators
/* Bitwiser
Operators */

import (
	"fmt"
)

func main() {
	var A = 20 // 0001 0100
	var B = 14 // 0000 1110

	fmt.Println(fmt.Sprintf("A is: %d and B is: %d", A, B))
	fmt.Println("A & B is: ", A&B)
	fmt.Println("A | B is: ", A|B)
	fmt.Println("A ^ B is: ", A^B)
	fmt.Println("Left Shift of A by 2 is: ", A<<2)
	fmt.Println("Right Shift of A by 2 is: ", A>>2)
}

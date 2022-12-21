package main

// Logical operators
import (
	"fmt"
)

func main() {
	A := false
	B := true

	fmt.Println(fmt.Sprintf("A is: %t and B is: %t", A, B))
	fmt.Println("A && B is: ", A && B)
	fmt.Println("A || B is: ", A || B)
	fmt.Println("Inverse of A is: ", !A)
	fmt.Println("Inverse of B is: ", !B)
}

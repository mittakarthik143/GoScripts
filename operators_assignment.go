package main

import (
	"fmt"
)

func main() {
	var A = 10
	B := 20
	var C int

	fmt.Println(fmt.Sprintf("A is: %d and B is: %d", A, B))
	A += B
	fmt.Println("A+=B is: ", A)
	A -= B
	fmt.Println("A-=B is: ", A)
	C = A + B
	fmt.Println("C(A+B) is: ", C)
	A *= B
	fmt.Println("A*=B is: ", A)
	A /= B
	fmt.Println("A/=B is: ", A)

}

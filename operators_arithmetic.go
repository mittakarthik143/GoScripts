// Arithmetic operators programme
package main

import "fmt"

func main() {
	var a, b int = 10, 20
	fmt.Printf("A is: %d and B is: %d = A+B is: %d \n", a, b, a+b)
	fmt.Printf("A is: %d and B is: %d = A-B is: %d \n", a, b, a-b)
	fmt.Printf("A is: %d and B is: %d = A*B is: %d \n", a, b, a*b)
	fmt.Printf("A is: %d and B is: %d = A/B is: %d \n", a, b, a/b)
	fmt.Println("B % A is: ", b%a)
}

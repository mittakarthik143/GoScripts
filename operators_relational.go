package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Println(fmt.Sprintf("A is: %d and B is: %d", a, b))

	fmt.Println("A<B is: ", a < b)
	fmt.Println("A>B is: ", a > b)
	fmt.Println("A<=B is: ", a <= b)
	fmt.Println("A>=B is: ", a >= b)
	fmt.Println("A==B is: ", a == b)
	fmt.Println("A!=B is: ", a != b)
}

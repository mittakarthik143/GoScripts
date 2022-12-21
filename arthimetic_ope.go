package main

import "fmt"

func main() {
	var a, b int = 10, 20

	fmt.Printf("addition is - %d\n", a+b)
	fmt.Printf("multiplication is - %d\n", a*b)
	fmt.Printf("substraction is - %d\n", a-b)
	fmt.Printf("division is - %d\n", a/b)
	fmt.Printf("modulo division is - %d\n", a%b)
	a++
	fmt.Printf("increment is - %d\n", a)
	a--
	fmt.Printf("decrement is - %d\n", a)
}

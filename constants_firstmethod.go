package main

import (
	"fmt"
)

func main() {
	const msg string = "Hello! world..."
	const number int = 100
	const checkbox bool = true
	const price float64 = 3.5

	fmt.Printf("Data is: %s and It's type is: %T \n", msg, msg)
	fmt.Printf("Data is: %d and It's type is: %T \n", number, number)
	fmt.Printf("Data is: %t and It's type is: %T \n", checkbox, checkbox)
	fmt.Printf("data is: %f and It's type is: %T \n", price, price)
}

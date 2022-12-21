package main

import "fmt"

// Constant variables once we declare after we can't change
// Var these are if we declare after that also we can change/modify the data
func main() {
	const msg = "Hello, World!"
	const number = 100
	const checkbox = true
	var price = 3.5
	fmt.Printf("Data is: %s and It's type is: %T \n", msg, msg)
	fmt.Printf("Data is: %d and It's type is: %T \n", number, number)
	fmt.Printf("Data is: %t and It's type is: %T \n", checkbox, checkbox)
	fmt.Printf("Data is: %f and It's type is: %T \n", price, price)
	price = 45.3
	fmt.Printf("Data is: %f and It's type is: %T \n", price, price)
}

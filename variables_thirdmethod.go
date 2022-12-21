package main

import "fmt"

func main() {
	string_data := "Go Lang"
	integer_data := 100
	bool_data := true
	float_data := 3.5

	fmt.Printf("Data is: %s and It's type is: %T \n", string_data, string_data)
	fmt.Printf("Data is: %d and It's type is: %T \n", integer_data, integer_data)
	fmt.Printf("Data is: %t and It's type is: %T \n", bool_data, bool_data)
	fmt.Printf("Data is: %f and It's type is: %T \n", float_data, float_data)
}

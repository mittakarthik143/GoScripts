package main

import (
	"fmt"
)

func main() {
	var string_data = "Go Lang"
	fmt.Printf("Data is: \"%s\" and it's type is: %T \n", string_data, string_data)

	var integer_data = 100
	fmt.Printf("Data is: %d and it's type is: %T \n", integer_data, integer_data)

	var bool_data = false
	fmt.Printf("Data is: \"%t\" and it's type is: %T \n", bool_data, bool_data)

	var float_data = 30.5
	fmt.Printf("Data is: %f and it's type is: %T \n", float_data, float_data)
}

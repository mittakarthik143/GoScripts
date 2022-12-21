package main

import (
	"fmt"
)

func main() {
	var string_data string
	string_data = "Go Lang"
	fmt.Printf("Data is: %s and It's type is: %T \n", string_data, string_data)

	var integer_data int
	integer_data = 100
	fmt.Printf("Data is: %d and It's type is: %T \n", integer_data, integer_data)

	var bool_data bool
	bool_data = true
	fmt.Printf("Data is: %t and It's type is: %T \n", bool_data, bool_data)

	var float_data float64
	float_data = 3.5
	fmt.Printf("Data is: %f and It's type is: %T \n", float_data, float_data)

}

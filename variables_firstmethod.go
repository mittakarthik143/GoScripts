package main

import "fmt"

func main() {
	var string_data string = "Go Lang"
	var integer_data int = 100
	var bool_data bool = true
	var float_data float64 = 34.89
	fmt.Printf("Data is: %s and it's type is: %T \n", string_data, string_data)
	fmt.Printf("Data is: %d and it's type is: %T \n", integer_data, integer_data)
	fmt.Printf("Data is: %t and it's type is: %T \n", bool_data, bool_data)
	fmt.Printf("Data is: %f and it's type is: %T \n", float_data, float_data)
}

package main

import "fmt"

func test_fun(val int) {
	// this function just normal it prints the value it won't return any information
	fmt.Printf("Value is: %d and It's type is: %T \n", val, val)
}

func test_fun_return(val int) int {
	// this function add 10 to val and return the data
	val += 10
	return val
}

func main() {
	test_fun(10)
	a := test_fun_return(10)
	fmt.Printf("after some of 10 to val is: %d \n", a)
}

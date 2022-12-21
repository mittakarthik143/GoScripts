package main

import (
	"fmt"
)

// nested if-else statement

func main() {
	a := 30

	if a < 100 {
		fmt.Println("A is smaller than 100")
		if a > 50 {
			fmt.Println("A is bigger than 50")
		} else {
			fmt.Println("A is smaller than 50")
		}
	} else {
		fmt.Println("A is bigger than 100")
		if a > 500 {
			fmt.Println("A is bigger than 500")
		} else {
			fmt.Println("A is smaller than 500")
		}
	}
}

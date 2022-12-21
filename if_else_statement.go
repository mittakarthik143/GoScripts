package main

// if statement

import (
	"fmt"
)

func main() {
	a := 101
	var info string
	if a < 100 {
		info = fmt.Sprintf("A is smaller than 100 and A value is: %d", a)
	} else {
		info = fmt.Sprintf("A is bigger than 100 and A value is: %d", a)
	}
	fmt.Println(info)
}

// Switch case methodology 1
package main

import "fmt"

func main() {
	var a int = 10
	switch a {
	case 1:
		fmt.Println("A value is '1'")
	case 2:
		fmt.Println("A value is '2'")
	case 3:
		fmt.Println("A value is '3'")
	case 4:
		fmt.Println("A value is '4'")
	case 5:
		fmt.Println("A value is '5'")
	case 10:
		fmt.Println("A value is '10'")
	default:
		fmt.Println("Not Found")
	}
}

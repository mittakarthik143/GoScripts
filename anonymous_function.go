package main

import "fmt"

func main() {
	func() {
		fmt.Println("It is a anonymous function")
	}()

	function := func() {
		fmt.Println("This is another anonymous function")
	}
	function()
}

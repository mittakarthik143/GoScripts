package main

import "fmt"

func main() {
	/*var num = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for index, value := range num {
		fmt.Printf("Index is: %d and Value is: %d \n", index, value)
		fmt.Println(index, "->", value)
	}*/
	var names = []string{"sai ram", "sai karthikeyan", "siva", "govinda"}
	for index, value := range names {
		fmt.Println(index, "->", value)
	}
}

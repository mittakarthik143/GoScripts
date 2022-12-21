package main

import (
	"fmt"
)

func main() {
	p := 10
	for i := 0; i < p; i++ {
		fmt.Println("i Value is: ", i)
		if i == 5 {
			break
		}
	}
	var name = []string{"sai ram", "sai karthikeyan", "balaji", "lakshmi", "hari", "naga"}
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
		if name[i] == "lakshmi" {
			break
		}
	}
}

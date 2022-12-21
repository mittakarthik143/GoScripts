package main

import (
	"fmt"
	"reflect"
)

func switch_fun(info interface{}) {
	q := reflect.TypeOf(info)
	fmt.Println(q)

	switch p := info.(type) {
	case string:
		fmt.Printf("Value is: %s and Type is: %T \n", p, p)
	case int:
		fmt.Printf("Value is: %d and Type is: %T \n", p, p)
	case float64:
		fmt.Printf("Value is: %f and Type is: %T \n", p, p)
	case bool:
		fmt.Printf("Value is: %t and Type is: %T \n", p, p)
	case rune:
		fmt.Printf("This value is rune %T\n", p)
	default:
		fmt.Println("Value is: ", p, "and Type is: ", q)
	}
}

func main() {
	switch_fun("Go Lang")
	switch_fun(10)
	switch_fun(3.5)
	switch_fun('q')
	switch_fun(true)
	switch_fun("")
}

package main

import "fmt"

func main(){
    /* Mixed Variables Declaration in GO */
    var a, b, c = 3, 2.0, "sai ram"
    fmt.Println("'a' value is - ", a)
    fmt.Printf("'a' type is - %T\n", a)
    fmt.Println("'b' value is - ", b)
    fmt.Printf("'b' type is - %T\n", b)
    fmt.Println("'c' value is - ", c)
    fmt.Printf("'c' type is - %T\n", c)
}
package main

import "fmt"

func main(){
    var a int = 10
    fmt.Printf(&a)
    fmt.Printf(*a)
}
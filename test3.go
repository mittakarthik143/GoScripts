package main

import "fmt"

func main(){
    /* Normal variable declaration */
    var x float64 = 20.0
    
    /* Dynamic variable declaration */
    y := 42
    
    fmt.Println("'x' value is -", x)
    fmt.Printf("'x' type of is - %T\n", x)
    
    fmt.Println("'y' value is -", y)
    fmt.Printf("'y' type of is - %T\n", y)
}
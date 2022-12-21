package main

import "fmt"

func main(){
    var a, b int = 10,20
    
    /* Equal */
    if (a==b){
        fmt.Println("a and b values are equal")
    } else {
        fmt.Println("a and b value are not equal")
    }
    
    /* not equal */
    if (a!=b){
        fmt.Println("a and b values are not equal")
    } else {
        fmt.Println("a and b values are equal")
    }
    
    /* greater */
    if (a>b){
        fmt.Println("a value is greater")
    } else {
        fmt.Println("b value is greater")
    }
    
    /* less */
    if (a<b){
        fmt.Println("b value is greater than a")
    } else {
        fmt.Println("a value is greater than b")
    }
    
    /* greater than and equal */
    if (a>=b){
        fmt.Println("a value is greater or equal")
    } else {
        fmt.Println("b value is greater or equal")
    }
}
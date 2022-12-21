package main

import "fmt"

func main(){
    var b int = 15
    var a int
    numbers := [6]int{1,2,3,5}/*[1,2,3,5,0,0]*/
    fmt.Println("b value is -",b)
    fmt.Println("numbers value is -",numbers)
    
    for a<b{
        fmt.Println("a value before increment - ",a)
        a++
        fmt.Println("a value after increment - ",a)
    }
    
    for a := 0; a <= 10; a++{
        fmt.Println("a value is - ",a)
    }
    
    for i, x := range numbers{
        fmt.Printf("x = %d value at %d\n", x,i+1)
    } 
}
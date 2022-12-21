package main

import "fmt"

func main(){
    var a int = 100
    var b int = 200
    
    if (a<b){
        fmt.Println("a is smaller")
        if (b<300){
            fmt.Println("a is bigger")
        } else {
            fmt.Println("300 is bigger")
        }
    } else {
        fmt.Println("b is smaller")
    }
}
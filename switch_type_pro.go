package main

import "fmt"

func main(){
    var x interface{}
    
    switch i := x.(type){
        case nil: fmt.Printf("type is %T\n",i)
        case int: fmt.Println("type is int")
        case float64: fmt.Println("type is float")
        case bool: fmt.Println("type is bool")
        case string: fmt.Println("type is string")
        default: fmt.Println("don't know the type")
    }
}
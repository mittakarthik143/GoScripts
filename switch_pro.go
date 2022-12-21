package main

import "fmt"

func main(){
    var a=90
    
    switch a{
        case 90: fmt.Println("Excellent")
        case 25: fmt.Println("Failed")
        case 60: fmt.Println("Third Class")
        case 75: fmt.Println("Second Class")
        case 85: fmt.Println("First Class")
    }
    
    var b = 70
    
    switch b{
        case 90: fmt.Println("Excellent")
        case 25: fmt.Println("Failed")
        case 70, 60: fmt.Println("Third Class")
        case 75: fmt.Println("Second Class")
        case 85: fmt.Println("First Class")
    }
}
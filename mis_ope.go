package main

import "fmt"

func main(){
    var a int = 4
    var b int32
    var c float64
    var ptr *int
    
    fmt.Printf("a type is - %T\n", a)
    fmt.Printf("b type is - %T\n", b)
    fmt.Printf("c type is - %T\n", c)
    
    /* example of & and * operator */
    ptr = &a  /* 'ptr' contains now address of a */
    fmt.Printf("a value is - %d\n",a)
    fmt.Printf("a addres is - %d\n",ptr)
    fmt.Printf("*ptr is %d\n", *ptr)
    
}
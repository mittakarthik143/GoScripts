package main

import "fmt"

func main(){
    var a int = 10
    var i int
    Sai: for i = 0; i<a ; i++{
        if i == 5{
            /*continue*/
            /*break*/
            goto Sai
        }else{
            i++
            fmt.Println(i)
        }
    }
}
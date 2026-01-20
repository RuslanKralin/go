package main

import "fmt"


func main() {
var num int = 10

pointer := &num
fmt.Println(pointer)
fmt.Println(*pointer)
}



func qwe(a *int) {
    *a = 20
}
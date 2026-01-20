package main

import "fmt"


func main() {

fmt.Println(res(5, 2), res(5, 3))
}



var res = func (a int, b int) int {
    return a * b
}
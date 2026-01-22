package main

import "fmt"

type User1 struct {
	Name  string
	Email string
	Age   int
}


func main() {
  number :=10
  var pointer = &number
  fmt.Println(pointer) // 0xc0000120c8
  var pointer2 *int

  fmt.Println(pointer2) // nil
 // fmt.Println("разыменование указателя",*pointer2) // будет ошибка, поэтому сначала проводят проверку

  if pointer2 != nil {
    fmt.Println("указатель обычный", *pointer2)
  } else {
    fmt.Println("указатель nil")
  }

}




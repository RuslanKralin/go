вывод типа данных 
fmt.Printf("Тип переменной text: %T\n", text)


func main() {
  var score int = 100
  asd := score == 5 // как условие
  println(asd)
}  // false

func main() {
 res := rand.Intn(10)
 for i := 0; i <= res; i++ {
    if i %2 == 0 {
        fmt.Println("четное", i)
        time.Sleep(500 * time.Millisecond)
    } else {
        fmt.Println("нечетное", i)
        time.Sleep(500 * time.Millisecond)
    }
}
} // тут вывод через 500 милисекунд 

бесконечный цикл
func main() {
for {
    fmt.Println("Hello")
    time.Sleep(1 * time.Second)
}
}
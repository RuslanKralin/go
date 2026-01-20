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

ОТЛОДЕННАЯ ФУНКЦИЯ (которая попажает в стек - структура данных)
defer func() {
    fmt.Println("defer")
}() // отрабоатет только тогда, когда завершиться функция, в рамках которой расположена функция defer, помним что они попадают в стек!

УКАЗАТЕЛИ - акпкменная которая хранит адрес в памяти
func main() {
var num int = 10

pointer := &num
fmt.Println(pointer) // 0xc000106008
fmt.Println(*pointer) // 10
}


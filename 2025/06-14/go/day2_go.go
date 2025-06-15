package main
import "fmt"

func main() {
    var age int = 25
    var price float64 = 19.99
    var grade rune = 'A'  // rune is used for characters
    var isStudent bool = true

    fmt.Println("Go")
    fmt.Println("Age:", age)
    fmt.Println("Price:", price)
    fmt.Println("Grade:", string(grade)) // convert rune to string for readable output
    fmt.Println("Is Student:", isStudent)
}

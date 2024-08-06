package main

import "fmt"

func main () {
    // Declaração de variáveis
    var a int = 10
    b := 20

    // Função
    result := add(a, b)
    fmt.Println("Resultado:", result)

    // Estrutura
    person := Person{Name: "Alice", Age: 25}
    fmt.Println(person)

    // Laço for
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // Condicional if
    if person.Age > 18 {
        fmt.Println("Maior de idade")
    } else {
        fmt.Println("Menor de idade")
    }

    // Slice
    slice := []int{1, 2, 3}
    slice = append(slice, 4)
    fmt.Println(slice)

    // Mapa
    m := make(map[string]int)
    m["key"] = 42
    fmt.Println(m)
}

func add(x, y int) int {
    return x + y
}

type Person struct {
    Name string
    Age  int
}
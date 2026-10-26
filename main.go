package main

import "fmt"

func calculate(a, b int) (int, int) {
    sum := a + b
    difference := a - b

    return sum, difference
}

func main() {
    sum, difference := calculate(20, 5)

    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", difference)
}
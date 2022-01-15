package main

import "fmt"

import "math"

func main() {
    fmt.Print("Enter value in foot: ")
    var input float64
    fmt.Scanf("%f", &input)           // округлим до 2х знаков в строке
    output := input * float64(0.3048) // точное значение
    rOutput := math.Round(output)     // округлим до целого
    sOutput := fmt.Sprintf("( %.2f)", output)
    fmt.Println("Value in Meters:", rOutput, sOutput )
}
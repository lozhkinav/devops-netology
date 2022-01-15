package main

import "fmt"

func main() {
    x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17}
    current := 0
    fmt.Println ("Список значений : ", x)
    for i, value := range x {
        if (i == 0) {
           current = value
        } else {
            if (value < current){
                current = value
            }
        }
    }
    fmt.Println("Минимальное число : ", current)
}
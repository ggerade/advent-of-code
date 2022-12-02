package main

import "fmt"

func determineMostCalories(elveCalories [][]int) int {
    most := 0
    for _, elf := range elveCalories {
        total := 0
        for _, calories := range elf {
             total += calories  
        }
        if total > most {
            most = total;
        }
    }
    return most;
}

func main() {
    elveCalories := fileTo2DimensionalSliceInts("challenge.txt");
    fmt.Println(determineMostCalories(elveCalories))
}

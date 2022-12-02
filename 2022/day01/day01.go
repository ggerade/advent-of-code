package main

import (
	"fmt"
	"sort"
)

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

func determineTotalOf3MostCalories(elveCalories [][]int) int {
    most := make([]int, 1)
    for _, elf := range elveCalories {
        total := 0
        for _, calories := range elf {
             total += calories  
        }
        most = append(most, total);
    }
    sort.Slice(most, func(i, j int) bool {
        return most[i] > most[j]
    })
    return most[0]+most[1]+most[2];
}

func main() {
    elveCalories := fileTo2DimensionalSliceInts("inputs/challenge.txt");
    most := determineMostCalories(elveCalories)
    fmt.Printf("Most calories: %d\n", most);
    most3 := determineTotalOf3MostCalories(elveCalories)
    fmt.Printf("Sum of 3 most calories: %d\n", most3);
}

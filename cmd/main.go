package main

import (
	"fmt"

	day1 "github.com/AbelXKassahun/AOC-2024/Day1"
)

func main() {
	// _, _, distance := day1.FirstHalf("../Day1/input.csv")
	// fmt.Println(distance)

	similarityScore := day1.SecondHalf("../Day1/input.csv")
	fmt.Println(similarityScore)
}
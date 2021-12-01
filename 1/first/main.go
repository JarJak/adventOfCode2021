package main

import (
	"adventOfCode/fileReader"
	"fmt"
	"math"
)

func main() {
	values, err := fileReader.ReadIntLines("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(countIncreases(values))
}

func countIncreases(values []int) int {
	prevValue := math.MaxInt
	counter := 0
	for _, value := range values {
		if value > prevValue {
			counter++
		}
		prevValue = value
	}

	return counter
}

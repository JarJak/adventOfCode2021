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
	var value int
	counter := 0
	for i, _ := range values {
		if i < 2 {
			continue
		}
		value = values[i-2] + values[i-1] + values[i]
		if value > prevValue {
			counter++
		}
		prevValue = value
	}

	return counter
}

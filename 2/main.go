package main

import (
	"adventOfCode/fileReader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	values, err := fileReader.ReadStringLines("input.txt")
	if err != nil {
		panic(err)
	}
	horizontal, depth := countPosition2(values)
	fmt.Println(horizontal * depth)
}

func countPosition(values []string) (int, int) {
	horizontal, depth := 0, 0
	for _, commandLine := range values {
		commandSplit := strings.Fields(commandLine)
		command := commandSplit[0]
		size, err := strconv.Atoi(commandSplit[1])
		if err != nil {
			panic(err)
		}
		switch command {
			case `forward`:
				horizontal += size
			case `up`:
				depth -= size
			case `down`:
				depth += size
		}
	}

	return horizontal, depth
}

func countPosition2(values []string) (int, int) {
	horizontal, depth, aim := 0, 0, 0
	for _, commandLine := range values {
		commandSplit := strings.Fields(commandLine)
		command := commandSplit[0]
		size, err := strconv.Atoi(commandSplit[1])
		if err != nil {
			panic(err)
		}
		switch command {
		case `forward`:
			horizontal += size
			depth += size * aim
		case `up`:
			aim -= size
		case `down`:
			aim += size
		}
	}

	return horizontal, depth
}

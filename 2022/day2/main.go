package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	println(match(input))

	println(match2(input))
}

func match(input string) int {

	transform := map[string]int{
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}
	count := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		count += transform[line]
	}
	return count
}

func match2(input string) int {
	transform := map[string]int{
		"A X": 3 + 0,
		"A Y": 1 + 3,
		"A Z": 2 + 6,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 2 + 0,
		"C Y": 3 + 3,
		"C Z": 1 + 6,
	}
	count := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		count += transform[line]
	}
	return count
}

package main

import _ "embed"

//go:embed input.txt
var input string

func main() {
	println(countFloor(input))
	println(countToBasement(input))
}

func countFloor(input string) int {
	floor := 0
	for _, char := range input {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return floor
}

func countToBasement(input string) int {
	pos := 0
	floor := 0
	for _, char := range input {
		pos++
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			return pos
		}
	}
	return -1
}

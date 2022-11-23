package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(countHouses(input))
	fmt.Println(countHousesWithRobo(input))
}

type location struct {
	x, y int
}

func countHouses(input string) int {
	visited := make(map[location]bool)
	visited[location{0, 0}] = true
	x, y := 0, 0
	for _, char := range input {
		switch char {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		visited[location{x, y}] = true
	}
	return len(visited)
}

func countHousesWithRobo(input string) int {
	visited := make(map[location]bool)
	visited[location{0, 0}] = true
	x, y := 0, 0
	x2, y2 := 0, 0
	for i, char := range input {
		if i%2 == 0 {
			switch char {
			case '^':
				y++
			case 'v':
				y--
			case '>':
				x++
			case '<':
				x--
			}
			visited[location{x, y}] = true
		} else {
			switch char {
			case '^':
				y2++
			case 'v':
				y2--
			case '>':
				x2++
			case '<':
				x2--
			}
			visited[location{x2, y2}] = true
		}
	}
	return len(visited)
}

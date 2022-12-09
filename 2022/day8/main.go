package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	testInput := `30373
25512
65332
33549
35390`
	tm := parseTreeMap(testInput)
	println(tm.countVisable())
	println(tm.findScenic())

	tm = parseTreeMap(input)
	println(tm.countVisable())
	println(tm.findScenic())
}

type treeMap struct {
	forward    [][]uint8
	transverse [][]uint8
}

func parseTreeMap(input string) treeMap {
	var tm treeMap
	lines := strings.Split(input, "\n")
	yMax := len(lines)
	xMax := len(lines[0])

	tm.forward = make([][]uint8, yMax)
	tm.transverse = make([][]uint8, xMax)
	for i := 0; i < yMax; i++ {
		tm.forward[i] = make([]uint8, xMax)
	}
	for i := 0; i < xMax; i++ {
		tm.transverse[i] = make([]uint8, yMax)
	}

	for y, line := range lines {
		for x, char := range line {
			tm.forward[y][x] = uint8(char)
			tm.transverse[x][y] = uint8(char)
		}
	}
	return tm
}

func (tm *treeMap) isVisable(x, y int) bool {
	if x == 0 || y == 0 || x == len(tm.forward[0]) || y == len(tm.forward) {
		return true
	}
	v := tm.forward[y][x]

	// look up
	if isMax(v, tm.transverse[x][:y]) {
		return true
	}
	// look down
	if isMax(v, tm.transverse[x][y+1:]) {
		return true
	}
	// look left
	if isMax(v, tm.forward[y][:x]) {
		return true
	}
	// look right
	if isMax(v, tm.forward[y][x+1:]) {
		return true
	}
	return false

}

func isMax(s uint8, slice []uint8) bool {
	for _, v := range slice {
		if v >= s {
			return false
		}
	}
	return true
}

func (tm *treeMap) countVisable() int {
	count := 0
	for y := range tm.forward {
		for x := range tm.forward[y] {
			if tm.isVisable(x, y) {
				count++
			}
		}
	}
	return count
}

func (tm *treeMap) findScenic() int {
	max := 0
	for y := range tm.forward {
		for x := range tm.forward[y] {
			score := tm.scenicScore(x, y)
			if score > max {
				max = score
			}
		}
	}
	return max
}

func (tm *treeMap) scenicScore(x, y int) int {
	var left, right, up, down int
	// left
	if x != 0 {
		left = countBackwardsUntil(tm.forward[y][x], tm.forward[y][:x])
	}
	// right
	if x != len(tm.forward[0]) {
		right = countUntil(tm.forward[y][x], tm.forward[y][x+1:])
	}
	// up
	if y != 0 {
		up = countBackwardsUntil(tm.transverse[x][y], tm.transverse[x][:y])
	}
	// down
	if y != len(tm.forward) {
		down = countUntil(tm.transverse[x][y], tm.transverse[x][y+1:])
	}

	return left * right * up * down
}

func countUntil(t uint8, slice []uint8) int {
	count := 0
	for _, v := range slice {
		count++
		if v >= t {
			return count
		}
	}
	return count
}
func countBackwardsUntil(t uint8, slice []uint8) int {
	count := 0
	for i := len(slice) - 1; i >= 0; i-- {
		count++
		if slice[i] >= t {
			return count
		}
	}
	return count
}

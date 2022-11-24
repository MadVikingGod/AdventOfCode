package main

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	g := grid{}
	insts, insts2 := parseInstructions(input)
	for _, i := range insts {
		g = i(g)
	}
	println(g.count())
	g = grid{}
	for _, i := range insts2 {
		g = i(g)
	}
	println(g.count())
}

var re = regexp.MustCompile(`(.+) (\d+),(\d+) through (\d+),(\d+)`)

func parseInstructions(input string) ([]instruction, []instruction) {
	var instructions []instruction
	var instructions2 []instruction
	for _, line := range strings.Split(input, "\n") {
		matches := re.FindStringSubmatch(line)
		x1, _ := strconv.Atoi(matches[2])
		y1, _ := strconv.Atoi(matches[3])
		x2, _ := strconv.Atoi(matches[4])
		y2, _ := strconv.Atoi(matches[5])
		var step1 step
		var step2 step
		switch matches[1] {
		case "turn on":
			step1 = turnOn
			step2 = turnUp
		case "turn off":
			step1 = turnOff
			step2 = turnDown
		case "toggle":
			step1 = toggle
			step2 = turnUp2
		}
		instructions = append(instructions, func(g grid) grid {
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					g[location{x, y}] = step1(g[location{x, y}])
				}
			}
			return g
		})
		instructions2 = append(instructions2, func(g grid) grid {
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					g[location{x, y}] = step2(g[location{x, y}])
				}
			}
			return g
		})
	}
	return instructions, instructions2
}

type instruction func(grid) grid

type step func(val int) int

type grid map[location]int

func (g grid) count() int {
	var count int
	for _, v := range g {
		count += v
	}
	return count
}

type location struct {
	x, y int
}

func turnOn(int) int {
	return 1
}

func turnOff(int) int {
	return 0
}
func toggle(val int) int {
	return 1 - val
}

func turnUp(val int) int {
	return val + 1
}
func turnDown(val int) int {
	if val > 0 {
		return val - 1
	}
	return 0
}
func turnUp2(val int) int {
	return val + 2
}

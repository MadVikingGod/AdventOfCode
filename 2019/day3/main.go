package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/madvikinggod/AdventOfCode/2018/helpers"
)

type location struct {
	x int
	y int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func (l location) distance(l2 location) int {

	return abs(l.x-l2.x) + abs(l.y-l2.y)
}
func (l location) add(l2 location) location {
	return location{l.x+l2.x, l.y+l2.y}
}

var dir = map[string]location{
	"R": location{1,0},
	"L": location{-1,0},
	"U": location{0,1},
	"D": location{0,-1},
}

type grid map[location]int

type direction struct {
	dir    string
	length int
}

func (g grid) addWire(steps []direction) {
	current := location{0, 0}

	count := 0
	for _, step := range steps {
		for i := 1; i <= step.length; i++ {
			count++
			current = current.add(dir[step.dir])

			if _, ok := g[current]; !ok {
				g[current] = count
			}

		}
	}
}
func (g grid) compare(g2 grid) (distance int, delay int) {
	distance = 99999999999
	delay =99999999999
	for loc := range g {
		if _, ok := g2[loc]; ok {
			dist := loc.distance(location{0, 0})
			if dist < distance {
				distance = dist
			}

			del := g[loc] + g2[loc]
			if del < delay {
				delay = del
			}
		}
	}
	return distance, delay
}

func main() {
	input, err := helpers.GetInput(3)
	if err != nil {
		panic(err)
	}

	grids := []grid{}
	for _, wire := range input {
		g := grid{}
		stepsRaw := strings.Split(wire, ",")
		steps := make([]direction, len(stepsRaw))
		for i, r := range stepsRaw {
			n, _ := strconv.Atoi(r[1:])
			steps[i].dir = string(r[0])
			steps[i].length = n
		}

		g.addWire(steps)

		grids = append(grids, g)
	}

	fmt.Println(grids[0].compare(grids[1]))

}

//43846 wrong

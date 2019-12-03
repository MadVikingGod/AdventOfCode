package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/madvikinggod/AdventOfCode/2018/helpers"
	"github.com/madvikinggod/AdventOfCode/2019/location"
)

type grid map[location.Location]int

type direction struct {
	dir    string
	length int
}

func (g grid) addWire(steps []direction) {
	current := location.New(0,0)
	var dir = map[string]location.Location{
		"R": location.New(1,0),
		"L": location.New(-1,0),
		"U": location.New(0,1),
		"D": location.New(0,-1),
	}
	// This is our step counter
	count := 0
	for _, step := range steps {
		for i := 1; i <= step.length; i++ {
			count++
			current = current.Add(dir[step.dir])

			// Because the counter is monotonic, if we have been here
			// the stored value will always be smallest
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
			dist := loc.Manhantan()
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

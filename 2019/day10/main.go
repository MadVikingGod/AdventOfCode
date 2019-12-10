package main

import (
	"fmt"
	"github.com/madvikinggod/AdventOfCode/2018/helpers"
	"github.com/madvikinggod/AdventOfCode/2019/location"
)

func main() {

	input, err := helpers.GetInput(10)
	if err != nil {
		panic(err)
	}

	space := map[location.Location]int{}

	for y, line := range input {
		for x, c := range line {
			if c == '#' {
				space[location.New(x, y)] = 0
			}
		}
	}
	fmt.Println(bestLocation(space))
	fmt.Println(space)

}

func bestLocation(space map[location.Location]int) (location.Location, int) {
	for base := range space {
		count := 0
		for astroid := range space {
			if astroid == base {
				continue
			}
			dir := base.Direction(astroid)
			for i := 1; i < 50; i++ {
				p := base.Add(dir.Mul(i))
				_, ok := space[p]
				if ok {
					if p == astroid {
						count++
					}
					break
				}
			}
		}

		space[base] = count
	}

	max := 0
	l := location.Location{}

	for loc, count := range space {
		if count > max {
			max = count
			l = loc
		}
	}

	return l, max

}

package main

import (
	"fmt"
	"github.com/madvikinggod/AdventOfCode/2018/helpers"
	"github.com/madvikinggod/AdventOfCode/2019/location"
	"sort"
)

type Space map[location.Location]int

func main() {

	input, err := helpers.GetInput(10)
	if err != nil {
		panic(err)
	}

	space := Space{}

	for y, line := range input {
		for x, c := range line {
			if c == '#' {
				space[location.New(x, y)] = 0
			}
		}
	}
	base, x := bestLocation(space)
	fmt.Println(base, x)
	//fmt.Println(space)

	count := 0
	for count < 200 {
		astroids := space.scan(base)
		if count+len(astroids) < 200 {
			count += len(astroids)
			for _, astroid := range astroids {
				delete(space, astroid)
			}
			continue
		}
		fmt.Println(count, astroids[199])

		break
	}

}

//1725

func (space Space) scan(base location.Location) []location.Location {
	astroids := []location.Location{}
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
					astroids = append(astroids, astroid)
				}
				break
			}
		}
	}
	less := func(i, j int) bool {
		I := base.Direction(astroids[i]).WeightedAngle()
		J := base.Direction(astroids[j]).WeightedAngle()
		return I < J
	}

	sort.Slice(astroids, less)
	return astroids
}

func bestLocation(space Space) (location.Location, int) {
	for base := range space {
		astroids := space.scan(base)

		space[base] = len(astroids)
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

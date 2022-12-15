package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

//go:embed testInput.txt
var testInput string

func main() {
	sensors := parseSensors(testInput)
	fmt.Println(CountFull(10, sensors) - 1)
	fmt.Println(findBeacon(20, sensors))

	sensors = parseSensors(input)
	fmt.Println(CountFull(2000000, sensors) - 1)
	//4793059 is too low
	p := findBeacon(4000000, sensors)
	fmt.Println(p, 4000000*p.x+p.y)

}

type point struct {
	x, y int
}

func (p point) distance(q point) int {
	return abs(p.x-q.x) + abs(p.y-q.y)
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type sensor struct {
	p        point
	distance int
}

func parseSensors(input string) []sensor {
	var sensors []sensor
	for _, line := range strings.Split(input, "\n") {
		s := sensor{}
		p2 := point{}
		//Sensor at x=3556832, y=3209801: closest beacon is at x=3520475, y=3164417
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.p.x, &s.p.y, &p2.x, &p2.y)
		s.distance = s.p.distance(p2)
		sensors = append(sensors, s)
	}
	return sensors
}

func CountFull(row int, sensors []sensor) int {
	count := 0
	maxX := 0
	minX := 100000000
	inRangeSensors := []sensor{}
	for _, s := range sensors {
		if s.p.x+s.distance > maxX {
			maxX = s.p.x + s.distance
		}
		if s.p.x-s.distance < minX {
			minX = s.p.x - s.distance
		}
		if s.p.y-s.distance <= row && s.p.y+s.distance >= row {
			inRangeSensors = append(inRangeSensors, s)
		}
	}

	for x := minX; x <= maxX; x++ {
		p := point{x, row}
		for _, s := range inRangeSensors {
			if p.distance(s.p) <= s.distance {
				count++
				break
			}
		}

	}

	return count
}

func findBeacon(maxXY int, sensors []sensor) point {
	for x := maxXY; x >= 0; x-- {
		for y := maxXY; y >= 0; y-- {
			p := point{x, y}
			covered := false
			for _, s := range sensors {
				if p.distance(s.p) <= s.distance {
					deltaX := abs(p.x - s.p.x)
					deltaY := s.distance - deltaX
					y = s.p.y - deltaY

					covered = true
					break
				}
			}
			if !covered {
				return p
			}
		}
	}
	return point{}
}

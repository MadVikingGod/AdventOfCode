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

	fmt.Println(ExternalSides(
		Cube(1, 1, 1),
		Cube(2, 1, 1),
	))

	c := parseCubes(testInput)
	fmt.Println(ExternalSides(c...))

	c = FillInternals(c)
	fmt.Println(ExternalSides(c...))

	c = parseCubes(input)
	fmt.Println(ExternalSides(c...))

	c = FillInternals(c)
	fmt.Println(ExternalSides(c...))

	//2354 is too low
}

func FillInternals(cubes []cube) []cube {
	found := map[cube]int{{0, 0, 0}: 1}
	for _, c := range cubes {
		found[c] = 2
	}
	for z := 0; z < 22; z++ {
		for y := 0; y < 22; y++ {
			for x := 0; x < 22; x++ {
				c := cube{x, y, z}
				if _, ok := found[c]; ok {
					continue
				}
				if AnyNeighborHas(found, c, 1) {
					found[c] = 1
					if AnyNeighborHas(found, c, 3) {
						found = Backfill(found, c)
					}
				} else {
					found[c] = 3
				}
			}
		}
	}

	out := []cube{}
	for c, v := range found {
		if v == 2 || v == 3 {
			out = append(out, c)
		}
	}
	return out
}

func Backfill(found map[cube]int, start cube) map[cube]int {
	queue := []cube{start}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		found[c] = 1
		for _, next := range getNeighbors(c) {
			if found[next] == 3 {
				queue = append(queue, next)
			}
		}
	}
	return found
}

func getNeighbors(c cube) []cube {
	return []cube{
		c.add(cube{1, 0, 0}),
		c.add(cube{-1, 0, 0}),
		c.add(cube{0, 1, 0}),
		c.add(cube{0, -1, 0}),
		c.add(cube{0, 0, 1}),
		c.add(cube{0, 0, -1}),
	}
}

func AnyNeighborHas(found map[cube]int, center cube, value int) bool {
	for _, n := range getNeighbors(center) {
		if found[n] == value {
			return true
		}
	}
	return false
}

func parseCubes(input string) []cube {
	lines := strings.Split(input, "\n")
	cubes := make([]cube, len(lines))
	for i, line := range lines {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		cubes[i] = Cube(x, y, z)
	}
	return cubes
}

func ExternalSides(cubes ...cube) int {
	sides := map[side]bool{}

	for _, c := range cubes {
		for _, s := range c.sides() {
			if _, ok := sides[s]; ok {
				delete(sides, s)
			} else {
				sides[s] = true
			}
		}
	}
	return len(sides)
}

type point struct {
	x, y, z int
}

func Cube(x, y, z int) cube {
	return cube{x, y, z}
}

type cube point

func (c cube) add(c2 cube) cube {
	return cube{c.x + c2.x, c.y + c2.y, c.z + c2.z}
}
func (c cube) sides() []side {

	p1, p2 := SideX(point(c))
	p3, p4 := SideY(point(c))
	p5, p6 := SideZ(point(c))

	return []side{p1, p2, p3, p4, p5, p6}
}

type side struct {
	p1, p2 point
}

func SideX(p point) (side, side) {
	return side{p, point{p.x, p.y + 1, p.z + 1}}, side{point{p.x + 1, p.y, p.z}, point{p.x + 1, p.y + 1, p.z + 1}}
}
func SideY(p point) (side, side) {
	return side{p, point{p.x + 1, p.y, p.z + 1}}, side{point{p.x, p.y + 1, p.z}, point{p.x + 1, p.y + 1, p.z + 1}}
}
func SideZ(p point) (side, side) {
	return side{p, point{p.x + 1, p.y + 1, p.z}}, side{point{p.x, p.y, p.z + 1}, point{p.x + 1, p.y + 1, p.z + 1}}
}

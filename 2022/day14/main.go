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
	c := parseCave(input)
	count := 0
	for p, done := c.fill(); !done; p, done = c.fill() {
		c.add(p, sand)
		// fmt.Println(c)
		count++
	}
	fmt.Println(count)

	c = parseCave(input)
	maxX := c.maxX + 200
	maxY := c.maxY + 2
	for x := c.minX - 200; x <= maxX; x++ {
		c.add(point{x, maxY}, rock)
	}
	count = 0
	var done bool
	var p point
	for !done {
		p, done = c.fill()
		c.add(p, sand)
		count++
	}
	fmt.Println(count - 1)
}

type cave struct {
	space map[point]fill

	minX, minY int
	maxX, maxY int
}

func parseCave(input string) *cave {
	c := &cave{
		space: map[point]fill{{500, 0}: filler},
		minX:  1000000,
		minY:  1000000,
	}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		p0 := parsePoint(parts[0])
		for i := 1; i < len(parts); i++ {
			p1 := parsePoint(parts[i])
			c.add(p0, rock)
			c.add(p1, rock)
			if p0.x == p1.x {
				for y := min(p0.y, p1.y); y <= max(p0.y, p1.y); y++ {
					c.add(point{p0.x, y}, rock)
				}
			} else {
				for x := min(p0.x, p1.x); x <= max(p1.x, p0.x); x++ {
					c.add(point{x, p0.y}, rock)
				}
			}
			p0 = p1
		}
	}
	return c
}

func (c *cave) add(p point, f fill) {
	if p.x < c.minX {
		c.minX = p.x
	}
	if p.y < c.minY {
		c.minY = p.y
	}
	if p.x > c.maxX {
		c.maxX = p.x
	}
	if p.y > c.maxY {
		c.maxY = p.y
	}
	c.space[p] = f
}

func (c *cave) fill() (point, bool) {
	p := point{500, 0}
	if c.space[p] == sand {
		return p, true
	}
	for {
		if p.y > c.maxY {
			return point{}, true
		}
		if c.space[p.add(down)] == air {
			p = p.add(down)
			continue
		}
		if c.space[p.add(downLeft)] == air {
			p = p.add(downLeft)
			continue
		}
		if c.space[p.add(downRight)] == air {
			p = p.add(downRight)
			continue
		}
		break
	}

	return p, false
}

func (c *cave) String() string {
	var sb strings.Builder
	for y := min(0, c.minY-1); y <= c.maxY+1; y++ {
		for x := c.minX - 1; x <= c.maxX+1; x++ {
			var b byte
			switch c.space[point{x, y}] {
			case air:
				b = '.'
			case rock:
				b = '#'
			case sand:
				b = 'o'
			case filler:
				b = '+'
			}
			sb.WriteByte(b)
		}
		sb.WriteByte('\n')
	}

	return sb.String()

}

type point struct {
	x, y int
}

func (p point) add(q point) point {
	return point{p.x + q.x, p.y + q.y}
}

var (
	down      = point{0, 1}
	downRight = point{1, 1}
	downLeft  = point{-1, 1}
)

type fill int

const (
	air fill = iota
	rock
	sand
	filler
)

func parsePoint(input string) point {
	var p point
	fmt.Sscanf(input, "%d,%d", &p.x, &p.y)
	return p
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

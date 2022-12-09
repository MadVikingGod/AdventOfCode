package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var testInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func main() {
	inst := parseInstructions(input)
	b := newBridge()
	for _, i := range inst {
		b.move(i)
	}
	fmt.Println(len(b.seenTail))
	fmt.Println(len(b.seenLongTail))

}

type point struct {
	x, y int
}

func (p point) add(p2 point) point {
	return point{p.x + p2.x, p.y + p2.y}
}
func (p point) sub(p2 point) point {
	return point{p.x - p2.x, p.y - p2.y}
}
func (p point) abs() point {
	if p.x < 0 {
		p.x = -p.x
	}
	if p.y < 0 {
		p.y = -p.y
	}
	return p
}
func (p point) dir() point {
	if p.x > 1 {
		p.x = 1
	}
	if p.x < -1 {
		p.x = -1
	}
	if p.y > 1 {
		p.y = 1
	}
	if p.y < -1 {
		p.y = -1
	}
	return p
}

func (p point) follow(leader point) point {
	d := leader.sub(p)
	if d.abs().x > 1 || d.abs().y > 1 {
		p = p.add(d.dir())
	}
	return p
}

var (
	left  = point{-1, 0}
	right = point{1, 0}
	up    = point{0, -1}
	down  = point{0, 1}
)

type bridge struct {
	knots []point

	seenTail     map[point]bool
	seenLongTail map[point]bool
}

func newBridge() *bridge {
	return &bridge{
		knots:        make([]point, 10),
		seenTail:     map[point]bool{{0, 0}: true},
		seenLongTail: map[point]bool{{0, 0}: true},
	}
}

func (b *bridge) move(p point) {
	b.knots[0] = b.knots[0].add(p)

	for i := 1; i < len(b.knots); i++ {
		b.knots[i] = b.knots[i].follow(b.knots[i-1])
	}
	b.seenTail[b.knots[1]] = true
	b.seenLongTail[b.knots[9]] = true
}

func parseInstructions(input string) []point {
	var instructions []point
	for _, s := range strings.Split(input, "\n") {
		if s == "" {
			continue
		}
		parts := strings.Split(s, " ")
		count, _ := strconv.Atoi(parts[1])
		for i := 0; i < count; i++ {
			switch parts[0] {
			case "U":
				instructions = append(instructions, up)
			case "D":
				instructions = append(instructions, down)
			case "L":
				instructions = append(instructions, left)
			case "R":
				instructions = append(instructions, right)
			}
		}
	}
	return instructions
}

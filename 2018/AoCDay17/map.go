package main

import (
	"fmt"
	"strings"
)

const (
	empty = iota
	clay
	spring
	stedyWater
	flowingWater
)

var strs = []string{".", "#", "+", "~", "|"}

type location struct {
	x int
	y int
}

type gameMap struct {
	squares map[location]int
	minY    int
	maxY    int
	minX    int
	maxX    int
}

func (m gameMap) below(loc location) int {
	loc.y++
	return m.squares[loc]
}
func (m gameMap) left(loc location) int {
	loc.x--
	return m.squares[loc]
}
func (m gameMap) right(loc location) int {
	loc.x++
	return m.squares[loc]
}

func (m gameMap) FindDown(loc location) (location, bool) {

	for m.below(loc) == empty && loc.y < m.maxY {
		loc.y++
	}

	return loc, !(loc.y < m.maxY)
}

func (m gameMap) FindLeft(loc location) (location, bool) {
	for m.left(loc) != clay && m.below(loc) != empty {
		loc.x--
	}
	return loc, m.below(loc) != empty
}
func (m gameMap) FindRight(loc location) (location, bool) {
	for m.right(loc) != clay && m.below(loc) != empty {
		loc.x++
	}
	return loc, m.below(loc) != empty
}

func (m *gameMap) FillDown(start, end location) {
	for start != end {
		m.squares[start] = flowingWater
		start.y++
	}
	m.squares[end] = flowingWater
}
func (m *gameMap) FillHorzStill(start, end location) {
	for start != end {
		m.squares[start] = stedyWater
		start.x++
	}
	m.squares[end] = stedyWater
}
func (m *gameMap) FillHorzFlowing(start, end location) {
	for start != end {
		m.squares[start] = flowingWater
		start.x++
	}
	m.squares[end] = flowingWater
}

func NewGameMap(input []string) *gameMap {
	gm := gameMap{
		squares: map[location]int{{500, 0}: spring},
		minY:    10000,
		maxY:    0,
		minX:    500,
		maxX:    500,
	}
	for _, line := range input {
		if line[0] == 'x' {
			var x, ymin, ymax int
			fmt.Sscanf(line, "x=%d, y=%d..%d", &x, &ymin, &ymax)
			for y := ymin; y <= ymax; y++ {
				gm.squares[location{x, y}] = clay
				if y < gm.minY {
					gm.minY = y
				}
				if y > gm.maxY {
					gm.maxY = y
				}
				if x < gm.minX {
					gm.minX = x
				}
				if x > gm.maxX {
					gm.maxX = x
				}
			}
		} else {
			var y, xmin, xmax int
			fmt.Sscanf(line, "y=%d, x=%d..%d", &y, &xmin, &xmax)
			for x := xmin; x <= xmax; x++ {
				gm.squares[location{x, y}] = clay
				if y < gm.minY {
					gm.minY = y
				}
				if y > gm.maxY {
					gm.maxY = y
				}
				if x < gm.minX {
					gm.minX = x
				}
				if x > gm.maxX {
					gm.maxX = x
				}
			}
		}

	}
	return &gm
}

func (m gameMap) String() string {
	b := strings.Builder{}
	for y := 0; y <= m.maxY; y++ {
		for x := m.minX - 1; x <= m.maxX+1; x++ {
			b.WriteString(strs[m.squares[location{x, y}]])
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (m gameMap) CountWater() int {
	n := 0
	for l, s := range m.squares {
		if l.y > m.maxY || l.y < m.minY {
			fmt.Println("gah", l)
		}
		if (s == flowingWater || s == stedyWater) && l.y <= m.maxY && l.y >= m.minY {
			n++
		}
	}
	return n
}

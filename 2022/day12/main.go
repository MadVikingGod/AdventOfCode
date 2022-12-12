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
	m := parseMap(input)
	fmt.Println(m.findPath())

	best := 1000000
	for y, row := range m.topo {
		for x, height := range row {
			if height == 0 {
				m.start = Point{x, y}
				d := m.findPath()
				if d != -1 && d < best {
					best = d
				}
			}
		}
	}
	fmt.Println(best)
}

type Map struct {
	topo       [][]int
	start, end Point
}

func (m *Map) getHeight(p Point) (int, error) {
	if p.y < 0 || p.y >= len(m.topo) {
		return 0, fmt.Errorf("y out of bounds: %d", p.y)
	}
	if p.x < 0 || p.x >= len(m.topo[p.y]) {
		return 0, fmt.Errorf("x out of bounds: %d", p.x)
	}
	return m.topo[p.y][p.x], nil
}

func (m *Map) findPath() int {
	type state struct {
		p Point
		d int
	}
	queue := []state{{m.start, 0}}
	visited := map[Point]bool{m.start: true}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		if s.p == m.end {
			return s.d
		}
		height := m.topo[s.p.y][s.p.x]
		for _, dir := range directions {
			next := s.p.Add(dir)
			if visited[next] {
				continue
			}
			nHeight, err := m.getHeight(next)
			if err != nil {
				continue
			}
			if nHeight > 1+height {
				continue
			}
			visited[next] = true
			queue = append(queue, state{next, s.d + 1})
		}
	}
	return -1
}

type Point struct {
	x, y int
}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

var directions = []Point{
	{0, -1}, // N
	{1, 0},  // E
	{0, 1},  // S
	{-1, 0}, // W
}

func parseMap(input string) *Map {
	var m Map
	for y, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		var row []int
		for x, c := range line {
			switch c {
			case 'S':
				row = append(row, 0)
				m.start = Point{x, y}
			case 'E':
				row = append(row, 'z'-'a')
				m.end = Point{x, y}
			default:
				row = append(row, int(c-'a'))
			}
		}
		m.topo = append(m.topo, row)
	}
	return &m
}

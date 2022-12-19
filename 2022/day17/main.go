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
	t := &tunnel{}
	wind := 0
	winds := testInput
	for i := 0; i < 1000000000000; i++ {
		current := shapes[i%len(shapes)]
		current.offset = point{2, len(t.fallen) + 3}
		next := current
		for {
			if winds[wind%len(winds)] == '<' {
				next = current.left()
			} else {
				next = current.right()
			}
			wind++
			if !t.collides(next) {
				current = next
			}

			next = current.down()
			if t.collides(next) {
				t.add(current)
				break
			} else {
				current = next
			}
		}
	}
	fmt.Println(len(t.fallen))
}

type tunnel struct {
	fallen [][7]bool
}

func (t *tunnel) collides(s shape) bool {
	for _, p := range s.Points() {
		if p.x < 0 || p.x >= 7 || p.y < 0 {
			return true
		}
		if p.y >= len(t.fallen) {
			continue
		}
		if t.fallen[p.y][p.x] {
			return true
		}
	}
	return false
}
func (t *tunnel) add(s shape) {
	for _, p := range s.Points() {
		for len(t.fallen) <= p.y {
			t.fallen = append(t.fallen, [7]bool{})
		}
		t.fallen[p.y][p.x] = true
	}
}

func (t *tunnel) String() string {
	buf := strings.Builder{}
	buf.WriteString("|.......|\n")
	for i := len(t.fallen) - 1; i >= 0; i-- {
		buf.WriteRune('|')
		for _, filled := range t.fallen[i] {
			if filled {
				buf.WriteRune('#')
			} else {
				buf.WriteRune('.')
			}
		}
		buf.WriteString("|\n")
	}
	buf.WriteString("+-------+")
	return buf.String()
}

type point struct {
	x, y int
}

func (p1 point) add(p2 point) point {
	return point{p1.x + p2.x, p1.y + p2.y}
}

type shape struct {
	offset point
	points []point
}

func (s shape) Points() []point {
	out := make([]point, len(s.points))
	for i, p := range s.points {
		out[i] = p.add(s.offset)
	}
	return out
}

func (s shape) left() shape {
	s.offset = s.offset.add(point{-1, 0})
	return s
}

func (s shape) right() shape {
	s.offset = s.offset.add(point{1, 0})
	return s
}
func (s shape) down() shape {
	s.offset = s.offset.add(point{0, -1})
	return s
}

var shapes = []shape{
	{
		points: []point{
			{0, 0}, {1, 0}, {2, 0}, {3, 0},
		},
	},
	{
		points: []point{
			{1, 2},
			{0, 1}, {1, 1}, {2, 1},
			{1, 0},
		},
	},
	{
		points: []point{
			{2, 2},
			{2, 1},
			{0, 0}, {1, 0}, {2, 0},
		},
	},
	{
		points: []point{
			{0, 3},
			{0, 2},
			{0, 1},
			{0, 0},
		},
	},
	{
		points: []point{
			{0, 1}, {1, 1},
			{0, 0}, {1, 0},
		},
	},
}

package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	current := new(input)
	nxt := &Ferry{seats: map[point]seat{}}

	for !current.Equal(nxt) {
		current, nxt = current.Next(nxt)
	}
	count := 0
	for _, s := range current.seats {
		if s == Full {
			count++
		}
	}
	fmt.Println(count)

	current = new(input)
	nxt = &Ferry{seats: map[point]seat{}}

	for !current.Equal(nxt) {
		current, nxt = current.Next2(nxt)
	}
	count = 0
	for _, s := range current.seats {
		if s == Full {
			count++
		}
	}
	fmt.Println(count)
}

type seat int

const (
	Floor = iota
	Empty
	Full
)

type point struct {
	x int
	y int
}

func (a point) Add(b point) point {
	return point{x: a.x + b.x, y: a.y + b.y}
}
func directions() []point {
	return []point{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
}

func (p point) neighbors() []point {
	pts := make([]point, 8)
	for i, dir := range directions() {
		pts[i] = p.Add(dir)
	}
	return pts
}

type Ferry struct {
	seats map[point]seat
	x, y  int
}

func new(s []string) *Ferry {
	f := &Ferry{
		y:     len(s),
		x:     len(s[0]),
		seats: map[point]seat{},
	}
	for y, row := range s {
		for x, seat := range row {
			switch seat {
			case '#':
				f.seats[point{x, y}] = Full
			case 'L':
				f.seats[point{x, y}] = Empty
			}
		}
	}
	return f
}

func (src *Ferry) Next(dst *Ferry) (*Ferry, *Ferry) {
	dst.x = src.x
	dst.y = src.y
	for p := range src.seats {
		dst.seats[p] = nextSeat(p, src)
	}
	return dst, src
}
func (src *Ferry) Next2(dst *Ferry) (*Ferry, *Ferry) {
	dst.x = src.x
	dst.y = src.y
	for p := range src.seats {
		dst.seats[p] = nextSeat2(p, src)
	}
	return dst, src
}

func (a *Ferry) Equal(b *Ferry) bool {
	return reflect.DeepEqual(a, b)
}
func (f *Ferry) farNeighbors(p point) []seat {
	seats := make([]seat, 8)
	for i, dir := range directions() {
		p := p.Add(dir)
		for f.seats[p] == Floor &&
			p.x >= 0 && p.x < f.x &&
			p.y >= 0 && p.y < f.y {
			p = p.Add(dir)
		}
		seats[i] = f.seats[p]
	}
	return seats
}

func nextSeat(p point, f *Ferry) seat {
	count := 0
	for _, n := range p.neighbors() {
		if f.seats[n] == Full {
			count++
		}
	}
	s := f.seats[p]
	return rule(s, count)
}
func nextSeat2(p point, f *Ferry) seat {
	count := 0
	for _, s := range f.farNeighbors(p) {
		if s == Full {
			count++
		}
	}
	s := f.seats[p]
	return rule2(s, count)

}

func rule(s seat, count int) seat {
	if s == Empty {
		if count == 0 {
			return Full
		}
		return Empty
	}
	if s == Full {
		if count > 3 {
			return Empty
		}
		return Full
	}
	return Floor
}
func rule2(s seat, count int) seat {
	if s == Empty {
		if count == 0 {
			return Full
		}
		return Empty
	}
	if s == Full {
		if count > 4 {
			return Empty
		}
		return Full
	}
	return Floor
}

func (f *Ferry) String() string {
	str := make([]string, f.y)
	for y := 0; y < f.y; y++ {
		s := make([]byte, f.x)
		for x := 0; x < f.x; x++ {
			switch f.seats[point{x, y}] {
			case Floor:
				s[x] = '.'
			case Empty:
				s[x] = 'L'
			case Full:
				s[x] = '#'
			}
		}
		str[y] = string(s)
	}
	return strings.Join(str, "\n")
}

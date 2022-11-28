package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

// var input = `.#.#.#
// ...##.
// #....#
// ..#...
// #.#..#
// ####..`

func main() {
	g := parse(input)
	for i := 0; i < 100; i++ {
		g = g.next()
	}
	println(g.count())

	g = parse(input)
	g[loc{0, 0}] = 1
	g[loc{0, 99}] = 1
	g[loc{99, 0}] = 1
	g[loc{99, 99}] = 1
	for i := 0; i < 100; i++ {
		g = g.next2()
	}
	println(g.count())
}

type grid map[loc]int
type loc struct {
	x, y int
}

func (g grid) next() grid {
	next := make(grid)
	for l := range g {
		sum := 0
		for _, n := range l.neighbors() {
			sum += g[n]
		}
		if g[l] == 1 {
			if sum == 2 || sum == 3 {
				next[l] = 1
			} else {
				next[l] = 0
			}
		} else {
			if sum == 3 {
				next[l] = 1
			} else {
				next[l] = 0
			}
		}
	}
	return next
}

func (g grid) next2() grid {
	next := g.next()
	next[loc{0, 0}] = 1
	next[loc{0, 99}] = 1
	next[loc{99, 0}] = 1
	next[loc{99, 99}] = 1
	return next
}

func (l loc) neighbors() []loc {
	return []loc{
		{l.x - 1, l.y},
		{l.x + 1, l.y},
		{l.x, l.y - 1},
		{l.x, l.y + 1},
		{l.x - 1, l.y - 1},
		{l.x + 1, l.y + 1},
		{l.x - 1, l.y + 1},
		{l.x + 1, l.y - 1},
	}
}

func (g grid) count() int {
	count := 0
	for _, v := range g {
		count += v
	}
	return count
}

func parse(input string) grid {
	g := make(grid)
	for y, line := range strings.Split(input, "\n") {
		for x, c := range line {
			if c == '#' {
				g[loc{x, y}] = 1
			} else {
				g[loc{x, y}] = 0
			}
		}
	}
	return g
}

func (g grid) String() string {
	var b strings.Builder
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if g[loc{x, y}] == 1 {
				b.WriteRune('#')
			} else {
				b.WriteRune('.')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

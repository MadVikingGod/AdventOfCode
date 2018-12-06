package main

import (
	"fmt"
)

var maxDistance = 10000

func main() {

	fmt.Println(len(inputs))

	queue := []loc{}
	board := make([][]loc, 350)
	for i := range board {
		board[i] = make([]loc, 350)
	}

	inf := map[int]bool{}
	for i, l := range inputs {
		queue = append(queue, loc{
			x:        l[0],
			y:        l[1],
			source:   i + 1,
			distance: 0,
		})
	}

	for len(queue) > 0 {
		l := queue[0]
		queue = queue[1:]

		if l.x < 0 || l.y < 0 || l.x > 349 || l.y > 349 {
			inf[l.source] = true
			continue
		}
		b := board[l.x][l.y]
		if l.x != b.x {
			board[l.x][l.y] = l
			queue = append(queue, l.up())
			queue = append(queue, l.down())
			queue = append(queue, l.left())
			queue = append(queue, l.right())
			continue
		}
		if b.source == l.source {
			continue
		}
		if l.distance < b.distance {
			board[l.x][l.y] = l
			continue
		}
		if b.distance == l.distance {
			board[l.x][l.y] = loc{
				x:        l.x,
				y:        l.y,
				source:   0,
				distance: 0,
			}
		}

	}

	counts := map[int]int{}

	for _, row := range board {
		for _, space := range row {
			counts[space.source]++
		}
	}
	max := 0
	for k, v := range counts {
		if _, ok := inf[k]; !ok && v > max {
			max = v
		}
	}

	fmt.Println(len(inf))
	fmt.Println(counts)
	fmt.Println(max)

	board2 := make([][]loc, 350)
	for i := range board2 {
		board2[i] = make([]loc, 350)
	}

	locs := []loc{}
	for i, l := range inputs {
		locs = append(locs, loc{
			x:        l[0],
			y:        l[1],
			source:   i + 1,
			distance: 0,
		})
	}

	safe := []loc{}
	for i, row := range board2 {
		for j := range row {
			sum := 0
			currentLoc := loc{
				x: i,
				y: j,
			}
			for _, l := range locs {
				sum = sum + l.dist(currentLoc)
			}
			if sum < maxDistance {
				currentLoc.distance = sum
				safe = append(safe, currentLoc)
			}
		}
	}
	fmt.Println(len(safe))

}

type loc struct {
	x        int
	y        int
	source   int
	distance int
}

func (l loc) dist(l2 loc) int {
	return abs(l.x-l2.x) + abs(l.y-l2.y)
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (l loc) up() loc {
	return loc{
		x:        l.x,
		y:        l.y + 1,
		source:   l.source,
		distance: l.distance + 1,
	}
}
func (l loc) down() loc {
	return loc{
		x:        l.x,
		y:        l.y - 1,
		source:   l.source,
		distance: l.distance + 1,
	}
}
func (l loc) right() loc {
	return loc{
		x:        l.x + 1,
		y:        l.y,
		source:   l.source,
		distance: l.distance + 1,
	}
}
func (l loc) left() loc {
	return loc{
		x:        l.x - 1,
		y:        l.y,
		source:   l.source,
		distance: l.distance + 1,
	}
}

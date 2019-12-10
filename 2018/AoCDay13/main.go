package main

import (
	"bytes"
	"fmt"
	"github.com/madvikinggod/AdventOfCode/2018/helpers"
	"sort"
)

func main() {
	Track := track{}
	carts := []*cart{}
	input, err := helpers.GetInput(13)
	if err != nil {
		panic(err)
	}
	//	input := strings.Split(`/>-<\
	//|   |
	//| /<+-\
	//| | | v
	//\>+</ |
	//  |   ^
	//  \<->/`, "\n")
	//input := strings.Split(`/->-\
	//|   |  /----\
	//| /-+--+-\  |
	//| | |  | v  |
	//\-+-/  \-+--/
	// \------/  `, "\n")

	for y, line := range input {
		for x, char := range line {
			switch char {
			case '-', '|', '+', '\\', '/':
				Track[New(x, y)] = string(char)
			case '^':
				Track[New(x, y)] = "|"
				carts = append(carts, &cart{
					loc:       New(x, y),
					direction: UP,
				})
			case '>':
				Track[New(x, y)] = "-"
				carts = append(carts, &cart{
					loc:       New(x, y),
					direction: RIGHT,
				})
			case '<':
				Track[New(x, y)] = "-"
				carts = append(carts, &cart{
					loc:       New(x, y),
					direction: LEFT,
				})
			case 'v':
				Track[New(x, y)] = "|"
				carts = append(carts, &cart{
					loc:       New(x, y),
					direction: DOWN,
				})
			}
		}
	}

	less := func(i, j int) bool {
		ci := carts[i]
		cj := carts[j]
		if ci.loc.y != cj.loc.y {
			return ci.loc.y < cj.loc.y
		}
		return ci.loc.x < cj.loc.x
	}
	count := 0
	display(Track, carts)
	for len(carts) > 1 {
		count++
		fmt.Println(count)
		sort.Slice(carts, less)
		for i, cart := range carts {
			cart.move(Track)
			for j, c2 := range carts {
				if i == j {
					continue
				}
				if cart.loc == c2.loc {
					fmt.Println("CRASH: ", cart.loc)
					cart.crashed = true
					c2.crashed = true
				}
			}
		}
		tmp := carts[:0]
		for _, cart := range carts {
			if !cart.crashed {
				tmp = append(tmp, cart)
			}
		}
		carts = tmp

		//display(Track, carts)
	}
	if len(carts) > 0 {
		fmt.Println("Last Cart:", carts[0])
		carts[0].move(Track)
		fmt.Println("Last Cart:", carts[0])

	}

}

type track map[location]string

const (
	UP = iota
	LEFT
	DOWN
	RIGHT
)

var dirFunc = map[int]func(location) location{
	UP:    Up,
	LEFT:  Left,
	DOWN:  Down,
	RIGHT: Right,
}

var turn = map[string]map[int]int{
	"\\": {
		UP:    LEFT,
		RIGHT: DOWN,
	},
	"/": {
		UP:   RIGHT,
		LEFT: DOWN,
	},
	"|": {
		UP:   UP,
		DOWN: DOWN,
	},
	"-": {
		LEFT:  LEFT,
		RIGHT: RIGHT,
	},
}
var cross = map[int]map[int]int{
	LEFT: {
		0: DOWN,
		1: LEFT,
		2: RIGHT,
	},
	RIGHT: {
		0: UP,
		1: RIGHT,
		2: DOWN,
	},
	UP: {
		0: LEFT,
		1: UP,
		2: RIGHT,
	},
	DOWN: {
		0: RIGHT,
		1: DOWN,
		2: LEFT,
	},
}

type cart struct {
	loc       location
	direction int
	next      int
	crashed   bool
}

func (c *cart) move(track track) {
	if c.crashed {
		return
	}
	c.loc = dirFunc[c.direction](c.loc)
	if track[c.loc] == "+" {
		c.direction = cross[c.direction][c.next]
		c.next = (c.next + 1) % 3
		return
	}
	c.direction = turn[track[c.loc]][c.direction]
}

type location struct {
	x, y int
}

func New(x, y int) location {
	return location{x, y}
}
func Up(loc location) location {
	return location{loc.x, loc.y - 1}
}
func Down(loc location) location {
	return location{loc.x, loc.y + 1}
}
func Right(loc location) location {
	return location{loc.x + 1, loc.y}
}
func Left(loc location) location {
	return location{loc.x - 1, loc.y}
}

func display(track track, carts []*cart) {
	output := [][]byte{}
	for i := 0; i < 151; i++ {
		output = append(output, bytes.Repeat([]byte(" "), 200))
		//output=append(output, []byte("             "))
	}

	for loc, t := range track {
		output[loc.y][loc.x] = t[0]
	}
	for _, cart := range carts {
		var c byte
		switch cart.direction {
		case UP:
			c = '^'
		case DOWN:
			c = 'v'
		case LEFT:
			c = '<'
		case RIGHT:
			c = '>'
		}
		output[cart.loc.y][cart.loc.x] = c
	}
	fmt.Println(string(bytes.Join(output, []byte("\n"))))

}

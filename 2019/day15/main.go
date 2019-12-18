package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/buger/goterm"
	"github.com/madvikinggod/AdventOfCode/2019/intcode"
	"github.com/madvikinggod/AdventOfCode/2019/location"

	"io"
)

func main() {

	p1r, p1w := io.Pipe()
	p2r, p2w := io.Pipe()

	ic := intcode.New(input)
	ic.In = p1r
	ic.Out = p2w
	ic.Register()
	go ic.Run()

	loc := location.New(0, 0)

	scanner := bufio.NewScanner(p2r)

	type backLocation struct {
		loc  location.Location
		back int
	}

	path := []backLocation{
		{location.New(-1000, -10000), -1},
	}

	for loc != path[0].loc {

		back := path[len(path)-1].loc
		backstr := path[len(path)-1].back

		if ship[loc] == 4 {
			loc = back
			path = path[0 : len(path)-1]

			fmt.Fprintf(p1w, "%d\n", backstr)
			if !scanner.Scan() {
				panic("Couldn't read from intcode")
			}
			text := scanner.Text()
			if text != "1" {
				panic("didn't move back correctly")
			}
			draw(loc, backstr, 1)
			continue
		}

		nextLoc := loc.Add(movement[ship[loc]])
		if !canGo(nextLoc) {
			ship[loc] += 1
			continue
		}

		if nextLoc == back {
			ship[loc] += 1
			continue
		}
		send := ship[loc] + 1
		fmt.Fprintf(p1w, "%d\n", send)
		if !scanner.Scan() {
			panic("Couldn't read from intcode")
		}
		resp, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			panic(err)
		}
		if resp == 0 {
			ship[nextLoc] = 5
			nextLoc = loc
		}
		if resp == 1 {
			path = append(path, backLocation{loc, backMovement[ship[loc]]})
		}
		if resp == 2 {
			oxy = nextLoc
			length = len(path)
			path = append(path, backLocation{loc, backMovement[ship[loc]]})
			draw(loc, send, resp)
			break
		}
		ship[loc] += 1
		loc = nextLoc
		draw(loc, send, resp)
	}

	queue := []location.Location{oxy}
	ship[oxy] = 100
	max := 100

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		nextval := ship[current] + 1
		if ship[current] > max {
			max = ship[current]
		}
		for _, dir := range movement {
			next := current.Add(dir)
			if ship[next] > 0 && ship[next] < 5 {
				ship[next] = nextval
				queue = append(queue, next)
			}
		}
		draw(loc, 1, max)
	}

}

func canGo(nextLoc location.Location) bool {
	return ship[nextLoc] == 0
}

var ship = map[location.Location]int{}
var oxy location.Location
var length int

var movement = map[int]location.Location{
	0: up,
	3: right,
	1: down,
	2: left,
	//4 is been there
	//5 is a wall
	//6 is the oxygen system
}
var backMovement = map[int]int{
	0: 2,
	1: 1,
	2: 4,
	3: 3,
}

var up = location.New(0, -1)
var down = location.New(0, 1)
var left = location.New(-1, 0)
var right = location.New(1, 0)

func draw(loc location.Location, send, resp int) {
	sendLookup := map[int]string{
		1: "up   ",
		2: "down ",
		3: "left ",
		4: "right",
	}

	lookup := map[int]string{
		0: " ",
		1: ".",
		2: ".",
		3: ".",
		4: ".",
		5: "#",
		6: "O",
	}
	//offset
	dx := 40
	dy := 30

	goterm.Clear()
	goterm.MoveCursor(0, 0)
	goterm.Print("Sent: ", sendLookup[send], ", Got: ", resp)
	for l, v := range ship {
		x, y := l.XY()
		goterm.MoveCursor(x+dx, y+dy)
		if v <= 6 {
			char := lookup[v]
			goterm.Print(char)
		} else {
			goterm.Print("O")
		}

	}
	goterm.MoveCursor(dx, dy)
	goterm.Print("*")

	if oxy != location.New(0, 0) {
		x, y := loc.XY()
		goterm.MoveCursor(x+dx, y+dy)
		goterm.Print("O")
	}
	if length != 0 {
		goterm.MoveCursor(1, 2)
		goterm.Print("PATH ", length)
	}
	goterm.Flush()
}

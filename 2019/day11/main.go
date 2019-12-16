package main

import (
	"bufio"
	"fmt"
	"github.com/madvikinggod/AdventOfCode/2019/day11/scratch"
	"github.com/madvikinggod/AdventOfCode/2019/intcode"
	"github.com/madvikinggod/AdventOfCode/2019/location"
	"io"
	"strconv"
	"strings"
	"sync"
)

type logger struct {
	queue string
}

func (l *logger) Write(b []byte) (int, error) {
	if len(l.queue) == 0 {
		l.queue = string(b)
		return len(b), nil
	}
	fmt.Printf("Logger: %+q, %+q\n", l.queue, string(b))
	l.queue = ""
	return len(b), nil
}

func main() {

	p1r, p1w := io.Pipe()
	p2r, p2w := io.Pipe()

	mw := io.MultiWriter(p2w, &logger{})

	ic := intcode.New(input)
	ic.In = p1r
	ic.Out = mw
	ic.Register()

	r := &robot{
		in:  p2r,
		out: p1w,
		dir: up,
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go r.run(wg)

	ic.Run()

	p1w.Close()
	p2r.Close()

	wg.Wait()
	fmt.Println(len(r.hull))
	fmt.Println(r.hull)
	whites := map[location.Location]int{}
	for loc, color := range r.hull {
		if color == 1 {
			whites[loc] = 1
		}
	}
	fmt.Println(whites)
}

type robot struct {
	in  io.Reader
	out io.Writer
	dir location.Location

	hull  map[location.Location]int
	count int
}

var c = []string{"Black", "White", "black"}
var t = []string{"Left", "Right"}

func (r *robot) run(wg *sync.WaitGroup) {
	r.out.Write([]byte("0\n"))
	scanner := bufio.NewScanner(r.in)

	loc := location.New(0, 0)
	for scanner.Scan() {
		color, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			panic(err)
		}
		if !scanner.Scan() {
			panic("Error reading direction")
		}
		dir, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		fmt.Printf("%d: At %v, paint %s, turn %s, looking %s - ", r.count, loc, c[color], t[dir], rev[r.dir])

		r.paint(loc, color)
		loc = r.move(dir, loc)
		fmt.Fprintf(r.out, "%d\n", r.hull[loc])

		fmt.Printf("moved to %v, it's color %s, now looking %s\n", loc, c[r.hull[loc]], rev[r.dir])
		scratch.Output(r.hull, loc, r.count)
		r.count++
	}
	scratch.WriteGif()
	wg.Done()
}

var up = location.New(0, 1)
var down = location.New(0, -1)
var left = location.New(-1, 0)
var right = location.New(1, 0)

var rev = map[location.Location]string{
	up:    "up",
	down:  "down",
	right: "right",
	left:  "left",
}

func (r *robot) paint(loc location.Location, color int) {
	r.hull[loc] = color
}

func (r *robot) move(i int, loc location.Location) location.Location {
	if i == 0 {
		switch r.dir {
		case up:
			r.dir = left
		case down:
			r.dir = right
		case left:
			r.dir = down
		case right:
			r.dir = up
		}
	} else if i == 1 {
		switch r.dir {
		case up:
			r.dir = right
		case down:
			r.dir = left
		case left:
			r.dir = up
		case right:
			r.dir = down
		}
	} else {
		panic(fmt.Sprintf("got invalid move %d, step %d", i, r.count))
	}
	return loc.Add(r.dir)
}

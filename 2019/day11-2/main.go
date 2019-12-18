package main

import (
	"bufio"
	"fmt"
	"github.com/buger/goterm"
	"github.com/madvikinggod/AdventOfCode/2019/intcode"
	"github.com/madvikinggod/AdventOfCode/2019/location"
	"io"
	"strconv"
	"sync"
)

func main() {
	p1r, p1w := io.Pipe()
	p2r, p2w := io.Pipe()

	r := &robot{
		wg: &sync.WaitGroup{},
	}
	r.wg.Add(1)
	go r.Run(p2r, p1w)

	ic := intcode.New(input)
	ic.In = p1r
	ic.Out = p2w
	ic.Register()
	ic.Run()
	p1w.Close()
	p2w.Close()

	r.wg.Wait()
	fmt.Println(len(hull))
}

var hull = map[location.Location]int{
	location.New(0, 0): 1,
}

var dir = map[int]location.Location{
	0: location.New(0, -1),
	1: location.New(1, 0),
	2: location.New(0, 1),
	3: location.New(-1, 0),
}

func (r *robot) Run(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	for {
		fmt.Fprintf(writer, "%d\n", hull[r.loc])
		if ok := scanner.Scan(); !ok {
			fmt.Println("No More input from IC")
			break
		}
		paint, _ := strconv.Atoi(scanner.Text())
		if paint != 0 && paint != 1 {
			fmt.Println("Got wrong paint", r.loc, paint)
			panic("")
		}
		if ok := scanner.Scan(); !ok {
			fmt.Println("No More input from IC")
			break
		}
		turn, _ := strconv.Atoi(scanner.Text())
		if turn != 0 && turn != 1 {
			fmt.Println("Got wrong paint", r.loc, paint)
			panic("")
		}

		hull[r.loc] = paint
		if turn == 0 {
			r.Left()
		}
		if turn == 1 {
			r.Right()
		}
		r.Move()
		draw(r, hull)
	}
	r.wg.Done()
}

type robot struct {
	loc       location.Location
	direction int
	wg        *sync.WaitGroup
}

func (r *robot) Right() {
	r.direction = (r.direction + 1) % 4
}
func (r *robot) Left() {
	r.direction -= 1
	if r.direction < 0 {
		r.direction = 3
	}
}
func (r robot) GetPaint(hull map[location.Location]int) int {
	return hull[r.loc]
}
func (r *robot) Move() {
	r.loc = r.loc.Add(dir[r.direction])
}

func draw(r *robot, hull map[location.Location]int) {
	dy := 50
	dx := 20
	lookup := map[int]string{
		0: "^",
		1: ">",
		2: "v",
		3: "<",
	}
	colorLookup := map[int]string{0: ".", 1: "#"}
	//goterm.Clear()
	for loc, color := range hull {
		x, y := loc.XY()
		goterm.MoveCursor(x+dx, y+dy)
		goterm.Print(colorLookup[color])
	}
	x, y := r.loc.XY()
	goterm.MoveCursor(x+dx, y+dy)
	goterm.Print(lookup[r.direction])
	goterm.Flush()
}

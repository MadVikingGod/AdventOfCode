package main

import (
	"bufio"
	"fmt"
	"github.com/buger/goterm"
	"github.com/madvikinggod/AdventOfCode/2019/intcode"
	"io"
	"strconv"
	"strings"
	"sync"
)

type reader string

func (r reader) Read(p []byte) (int, error) {
	return copy(p, []byte(r)), nil
}

func main() {
	r, w := io.Pipe()

	input[0] = 2

	ic := intcode.New(input)
	ic.In = reader("0\n")
	ic.Out = w
	ic.Register()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go draw(r, wg)

	ic.Run()
	w.Close()
	wg.Wait()

	count := 0
	for _, pix := range screen {
		if pix == 2 {
			count++
		}
	}
	fmt.Println(count)
}

type location struct {
	x, y int
}

var screen = map[location]int{}

func draw(in io.Reader, wg *sync.WaitGroup) {
	scan := bufio.NewScanner(in)

	state := 0
	x := 0
	y := 0

	maxX := 0
	for scan.Scan() {
		text := scan.Text()
		numb, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil {
			fmt.Println("Error reading text,", text, err)
			return
		}
		switch state {
		case 0:
			if numb > maxX {
				maxX = numb
			}
			x = numb

		case 1:
			y = numb
		case 2:
			screen[location{x, y}] = numb
			if numb == 3 {
				for i := 0; i < maxX; i++ {
					screen[location{i, y}] = 3
				}
			}
			if x == 1 && y == 0 {
				fmt.Println("Score: ", numb)
			}
			DrawScreen(maxX)
		}
		state = (state + 1) % 3
	}

	wg.Done()
}

func DrawScreen(max int) {
	lookup := map[int]string{
		0: " ",
		1: "x",
		2: "#",
		3: "-",
		4: "@",
	}
	goterm.Clear()
	for loc, char := range screen {
		goterm.MoveCursor(loc.x+1, loc.y+1)
		goterm.Print(lookup[char])
	}
	goterm.MoveCursor(40, 1)
	goterm.Print("Max X: ", max)
	goterm.MoveCursor(40, 3)
	goterm.Print("Score: ", screen[location{-1, 0}])
	goterm.Flush()
}

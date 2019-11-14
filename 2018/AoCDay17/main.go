package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	buf, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputs := strings.Split(string(buf), "\n")

	gm := NewGameMap(inputs)
	// fmt.Print(gm)

	getStedyState(gm)

	fmt.Print(gm)
	fmt.Println(gm.CountWater())
}

func getStedyState(gm *gameMap) {
	queue := []location{location{500, 1}}
	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]

		if gm.squares[start] == stedyWater {
			continue
		}
		if gm.squares[start] == flowingWater && gm.below(start) == flowingWater {
			continue
		}

		end, isinf := gm.FindDown(start)
		gm.FillDown(start, end)
		if isinf {
			continue
		}
		if gm.below(end) == flowingWater {
			continue
		}
	loop:
		l, lclose := gm.FindLeft(end)
		r, rclose := gm.FindRight(end)
		if lclose && rclose {
			gm.FillHorzStill(l, r)
			end.y--
			goto loop
		}

		gm.FillHorzFlowing(l, r)

		if !lclose {
			l.y++
			queue = append(queue, l)
		}
		if !rclose {
			r.y++
			queue = append(queue, r)
		}
	}
}

package main

import (
	"fmt"
	"github.com/madvikinggod/AdventOfCode/2019/intcode"
	"strconv"
	"strings"
)

var lookup = map[string]string{
	"46":  ".",
	"35":  "#",
	"60":  "<",
	"94":  "^",
	"62":  ">",
	"118": "v",
}

func main() {
	output := &strings.Builder{}

	ic := intcode.New(input)
	ic.Out = output
	ic.In = buildProgram(`A,B,B,C,A
L,12,R,8
L,6,R,8
L,12
n
`)
	ic.Register()
	ic.Run()

	out := strings.Split(output.String(), "10\n")
	fmt.Println("dust: ", output.String()[len(output.String())-10:len(output.String())-1])
	scafold := make([][]string, len(out))
	for i, s := range out {
		temp := strings.Split(s, "\n")
		for j, x := range temp {
			temp[j] = lookup[x]
		}
		scafold[i] = temp

	}
	MAXY = len(scafold)
	MAXX = len(scafold[0])

	for _, s := range scafold {
		fmt.Println(strings.Join(s, ""))
	}

	//sum :=0
	//for y,row := range scafold {
	//	for x := range row {
	//		if isCross(x,y, scafold) {
	//			sum+=x*y
	//			fmt.Println(x,y, x*y)
	//		}
	//	}
	//}
	//fmt.Println(sum)
	//6834 tool high

}

var MAXX = 0
var MAXY = 0

func OOB(x, y int) error {
	return fmt.Errorf("out of bounds, (%d,%d)", x, y)
}
func left(x, y int) (int, int, error) {
	if x == 0 {
		return 0, 0, OOB(x-1, y)
	}
	return x - 1, y, nil
}
func right(x, y int) (int, int, error) {
	if x == MAXX {
		return 0, 0, OOB(x+1, y)
	}
	return x + 1, y, nil
}
func up(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, OOB(x, y-1)
	}
	return x, y - 1, nil
}
func down(x, y int) (int, int, error) {
	if x == MAXX {
		return 0, 0, OOB(x, y+1)
	}
	return x, y + 1, nil
}

func isCross(x, y int, scaffold [][]string) bool {
	if x < 0 || y < 0 || x >= MAXX || y >= MAXY {
		return false
	}
	if scaffold[y][x] != "#" {
		return false
	}
	check := []func(int, int) (int, int, error){
		left,
		right,
		up,
		down,
	}

	for _, f := range check {
		x, y, err := f(x, y)
		if err != nil {
			return false
		}
		if scaffold[y][x] != "#" {
			return false
		}
	}
	return true
}
func buildProgram(program string) *strings.Reader {
	out := strings.Builder{}
	for _, char := range program {
		out.WriteString(strconv.Itoa(int(char)))
		out.WriteString("\n")
	}
	return strings.NewReader(out.String())
}

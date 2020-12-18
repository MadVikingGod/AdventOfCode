package main

import (
	"fmt"
	"strings"
)

func main() {

	pocket := New3d(input)
	pocket = pocket.Step() //step1
	pocket = pocket.Step() //step2
	pocket = pocket.Step() //step3
	pocket = pocket.Step() //step4
	pocket = pocket.Step() //step5
	pocket = pocket.Step() //step6

	fmt.Println(pocket.Actives())

	pocket = New4d(input)
	pocket = pocket.Step() //step1
	pocket = pocket.Step() //step2
	pocket = pocket.Step() //step3
	pocket = pocket.Step() //step4
	pocket = pocket.Step() //step5
	pocket = pocket.Step() //step6

	fmt.Println(pocket.Actives())
}

type point interface {
	Min(point) point
	Max(point) point
	Neighbors() []point
	Range(max point) []point
	AddScaler(int) point
}

type dim struct {
	min, max point
	space    map[point]int
}

func New3d(s []string) dim {
	min := orig3d
	max := point3d{x: len(s[0]), y: len(s), z: 0}
	space := map[point]int{}
	for y, line := range s {
		for x, char := range line {
			if char == '#' {
				space[point3d{x, y, 0}] = 1
			}
		}
	}
	return dim{
		min:   min,
		max:   max,
		space: space,
	}
}

func New4d(s []string) dim {
	min := orig4d
	max := point4d{x: len(s[0]), y: len(s), z: 0}
	space := map[point]int{}
	for y, line := range s {
		for x, char := range line {
			if char == '#' {
				space[point4d{x, y, 0, 0}] = 1
			}
		}
	}
	return dim{
		min:   min,
		max:   max,
		space: space,
	}
}
func (d dim) Count(p point) int {
	count := 0
	for _, b := range p.Neighbors() {
		count += d.space[b]
	}
	return count
}

func (d dim) Step() dim {
	min := d.min
	max := d.max
	space := map[point]int{}
	for _, p := range min.AddScaler(-1).Range(max.AddScaler(1)) {
		count := d.Count(p)
		if d.space[p] == 1 && (count == 2 || count == 3) {
			space[p] = 1
			continue
		}
		if d.space[p] == 0 && count == 3 {
			space[p] = 1
			min = min.Min(p)
			max = max.Max(p)
			continue
		}

	}
	return dim{
		min:   min,
		max:   max,
		space: space,
	}
}

func (d dim) Actives() int {
	count := 0
	for _, cell := range d.space {
		count += cell
	}
	return count
}

func (d dim) String() string {
	if _, ok := d.max.(point3d); ok {
		return d.String3d()
	}
	if _, ok := d.max.(point4d); ok {
		return d.String4d()
	}
	return "¯\\_(ツ)_/¯"
}

func (d dim) String3d() string {
	min := d.min.(point3d)
	max := d.max.(point3d)
	out := &strings.Builder{}
	for z := min.z; z < max.z+1; z++ {
		fmt.Fprintf(out, "z=%d\n", z)
		for y := min.y; y < max.y+1; y++ {
			for x := min.x; x < max.x+1; x++ {
				if d.space[point3d{x, y, z}] == 1 {
					out.WriteRune('#')
				} else {
					out.WriteRune('.')
				}
			}
			fmt.Fprintln(out, "")
		}
	}
	return out.String()

}
func (d dim) String4d() string {
	min := d.min.(point4d)
	max := d.max.(point4d)
	out := &strings.Builder{}
	for w := min.w; w < max.w+1; w++ {
		for z := min.z; z < max.z+1; z++ {
			fmt.Fprintf(out, "z=%d, w=%d\n", z, w)
			for y := min.y; y < max.y+1; y++ {
				for x := min.x; x < max.x+1; x++ {
					if d.space[point4d{x, y, z, w}] == 1 {
						out.WriteRune('#')
					} else {
						out.WriteRune('.')
					}
				}
				fmt.Fprintln(out, "")
			}
		}
	}
	return out.String()

}

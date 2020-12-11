package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"reflect"
	"strings"

	"github.com/nfnt/resize"
)

func main() {
	current := new(input)
	nxt := &Ferry{seats: map[point]seat{}}

	os.MkdirAll("2020/day11/img/part1", 0755)
	i := 0
	for !current.Equal(nxt) {
		f, _ := os.Create(fmt.Sprintf("2020/day11/img/part1/img%03d.png", i))
		png.Encode(f, current.Img())
		current, nxt = current.Next(nxt)
		f.Close()
		i++
	}
	count := 0
	for _, s := range current.seats {
		if s == Full {
			count++
		}
	}
	fmt.Println(count)

	os.MkdirAll("2020/day11/img/part2", 0755)
	current = new(input)
	nxt = &Ferry{seats: map[point]seat{}}

	i = 0
	for !current.Equal(nxt) {
		f, _ := os.Create(fmt.Sprintf("2020/day11/img/part2/img%03d.png", i))
		png.Encode(f, current.Img())
		current, nxt = current.Next2(nxt)
		f.Close()
		i++
	}
	count = 0
	for _, s := range current.seats {
		if s == Full {
			count++
		}
	}
	fmt.Println(count)
}

type seat int

const (
	Floor = iota
	Empty
	Full
)

type point struct {
	x int
	y int
}

func (a point) Add(b point) point {
	return point{x: a.x + b.x, y: a.y + b.y}
}
func directions() []point {
	return []point{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
}

func (p point) neighbors() []point {
	pts := make([]point, 8)
	for i, dir := range directions() {
		pts[i] = p.Add(dir)
	}
	return pts
}

type Ferry struct {
	seats map[point]seat
	x, y  int
}

func new(s []string) *Ferry {
	f := &Ferry{
		y:     len(s),
		x:     len(s[0]),
		seats: map[point]seat{},
	}
	for y, row := range s {
		for x, seat := range row {
			switch seat {
			case '#':
				f.seats[point{x, y}] = Full
			case 'L':
				f.seats[point{x, y}] = Empty
			}
		}
	}
	return f
}

func (src *Ferry) Next(dst *Ferry) (*Ferry, *Ferry) {
	dst.x = src.x
	dst.y = src.y
	for p := range src.seats {
		dst.seats[p] = nextSeat(p, src)
	}
	return dst, src
}
func (src *Ferry) Next2(dst *Ferry) (*Ferry, *Ferry) {
	dst.x = src.x
	dst.y = src.y
	for p := range src.seats {
		dst.seats[p] = nextSeat2(p, src)
	}
	return dst, src
}

func (a *Ferry) Equal(b *Ferry) bool {
	return reflect.DeepEqual(a, b)
}
func (f *Ferry) farNeighbors(p point) []seat {
	seats := make([]seat, 8)
	for i, dir := range directions() {
		p := p.Add(dir)
		for f.seats[p] == Floor &&
			p.x >= 0 && p.x < f.x &&
			p.y >= 0 && p.y < f.y {
			p = p.Add(dir)
		}
		seats[i] = f.seats[p]
	}
	return seats
}

func nextSeat(p point, f *Ferry) seat {
	count := 0
	for _, n := range p.neighbors() {
		if f.seats[n] == Full {
			count++
		}
	}
	s := f.seats[p]
	return rule(s, count)
}
func nextSeat2(p point, f *Ferry) seat {
	count := 0
	for _, s := range f.farNeighbors(p) {
		if s == Full {
			count++
		}
	}
	s := f.seats[p]
	return rule2(s, count)

}

func rule(s seat, count int) seat {
	if s == Empty {
		if count == 0 {
			return Full
		}
		return Empty
	}
	if s == Full {
		if count > 3 {
			return Empty
		}
		return Full
	}
	return Floor
}
func rule2(s seat, count int) seat {
	if s == Empty {
		if count == 0 {
			return Full
		}
		return Empty
	}
	if s == Full {
		if count > 4 {
			return Empty
		}
		return Full
	}
	return Floor
}

func (f *Ferry) String() string {
	str := make([]string, f.y)
	for y := 0; y < f.y; y++ {
		s := make([]byte, f.x)
		for x := 0; x < f.x; x++ {
			switch f.seats[point{x, y}] {
			case Floor:
				s[x] = '.'
			case Empty:
				s[x] = 'L'
			case Full:
				s[x] = '#'
			}
		}
		str[y] = string(s)
	}
	return strings.Join(str, "\n")
}

var (
	colors = map[seat]color.NRGBA{
		Floor: {R: 0, G: 0x33, B: 0x42, A: 255},
		Empty: {R: 0, G: 0xD1, B: 0x99, A: 255},
		Full:  {R: 0xf6, G: 0x3f, B: 0x64, A: 255},
	}
)

func (f *Ferry) Img() image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, f.x-1, f.y-1))
	for p, s := range f.seats {
		img.Set(p.x, p.y, colors[s])
	}

	out := resize.Resize(uint(f.x*10), uint(f.y*10), img, resize.NearestNeighbor)
	return out
}

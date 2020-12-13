package main

import (
	"fmt"
	"strconv"
)

var funcs = map[byte]func(int) func(*ship){
	'N': North,
	'S': South,
	'E': East,
	'W': West,
	'L': Left,
	'R': Right,
	'F': Forward,
}

func main() {

	s := new()
	for _, in := range input {
		x, _ := strconv.Atoi(in[1:])
		funcs[in[0]](x)(s)

	}
	fmt.Println(s.location.Distance(point{}))
	fmt.Println(s.location2.Distance(point{}))

	// part2 63078 too high
}

type ship struct {
	dir      point
	location point

	waypoint  point
	location2 point
}

func new() *ship {
	return &ship{east, point{}, point{10, 1}, point{}}
}

func North(i int) func(*ship) {
	return func(s *ship) {
		s.location = s.location.Add(north.Mul(i))
		s.waypoint = s.waypoint.Add(north.Mul(i))
	}
}
func South(i int) func(*ship) {
	return func(s *ship) {
		s.location = s.location.Add(south.Mul(i))
		s.waypoint = s.waypoint.Add(south.Mul(i))
	}
}
func East(i int) func(*ship) {
	return func(s *ship) {
		s.location = s.location.Add(east.Mul(i))
		s.waypoint = s.waypoint.Add(east.Mul(i))
	}
}
func West(i int) func(*ship) {
	return func(s *ship) {
		s.location = s.location.Add(west.Mul(i))
		s.waypoint = s.waypoint.Add(west.Mul(i))
	}
}
func Forward(i int) func(*ship) {
	return func(s *ship) {
		s.location = s.location.Add(s.dir.Mul(i))
		s.location2 = s.location2.Add(s.waypoint.Mul(i))
	}
}
func Left(d int) func(*ship) {
	sin := map[int]int{
		0:   0,
		90:  1,
		180: 0,
		270: -1,
	}[d]
	cos := map[int]int{
		0:   1,
		90:  0,
		180: -1,
		270: 0,
	}[d]
	return func(s *ship) {
		s.dir.x, s.dir.y = s.dir.x*cos-s.dir.y*sin, s.dir.x*sin+s.dir.y*cos
		s.waypoint.x, s.waypoint.y = s.waypoint.x*cos-s.waypoint.y*sin, s.waypoint.x*sin+s.waypoint.y*cos
	}
}
func Right(d int) func(*ship) {
	return Left(map[int]int{
		90:  270,
		180: 180,
		270: 90,
	}[d])

}

type point struct {
	x int
	y int
}

func (a point) Add(b point) point {
	return point{x: a.x + b.x, y: a.y + b.y}
}
func (a point) Mul(b int) point {
	return point{a.x * b, a.y * b}
}
func (a point) Distance(b point) int {
	x := a.x - b.x
	if x < 0 {
		x = -x
	}
	y := a.y - b.y
	if y < 0 {
		y = -y
	}
	return x + y
}

var (
	north = point{0, 1}
	south = point{0, -1}
	east  = point{1, 0}
	west  = point{-1, 0}
)

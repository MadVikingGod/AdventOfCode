package location

import (
	"image"
	"math"
)

type Location struct {
	x int
	y int
}

func New(x, y int) Location {
	return Location{x, y}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func gcd(a, b int) int {
	a = abs(a)
	b = abs(b)

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (l Location) Distance(l2 Location) int {
	return abs(l.x-l2.x) + abs(l.y-l2.y)
}

func (l Location) Manhantan() int {
	return abs(l.x) + abs(l.y)
}

func (l Location) Add(l2 Location) Location {
	return Location{l.x + l2.x, l.y + l2.y}
}

func (l Location) Direction(l2 Location) Location {
	x := l2.x - l.x
	y := l2.y - l.y
	d := gcd(x, y)

	return Location{x / d, y / d}
}

func (l Location) Mul(a int) Location {
	return Location{l.x * a, l.y * a}
}

//WeightedAngle This isn't a real angle, but the ratio of opp/adj
// It expects the abs(ratio) to be less then 50.
// It is weighted this way to make up (0,-1) the lowest, and
func (dir Location) WeightedAngle() float64 {
	if dir.x == 0 && dir.y < 0 {
		return -100.0
	}
	if dir.x == 0 && dir.y > 0 {
		return 0
	}
	if dir.x > 0 {
		return float64(dir.y)/float64(dir.x) - 50
	}
	return float64(dir.y)/float64(dir.x) + 50
}

func (l Location) Angle(l2 Location) float64 {
	x := float64(l2.x - l.x)
	y := float64(l2.y - l.y)
	return math.Atan2(y, x)
}

func (l Location) Max(l2 Location) Location {
	x, y := 0, 0
	if l.x > l2.x {
		x = l.x
	} else {
		x = l2.x
	}
	if l.y > l2.y {
		y = l.y
	} else {
		y = l2.y
	}
	return New(x, y)
}
func (l Location) Min(l2 Location) Location {
	x, y := 0, 0
	if l.x < l2.x {
		x = l.x
	} else {
		x = l2.x
	}
	if l.y < l2.y {
		y = l.y
	} else {
		y = l2.y
	}
	return New(x, y)
}

func (l Location) Point() image.Point {
	return image.Pt(l.x, l.y)
}
func Rect(l1, l2 Location) image.Rectangle {
	return image.Rect(l1.x, l1.y, l2.x, l2.y)
}

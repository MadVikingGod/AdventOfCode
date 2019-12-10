package location

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

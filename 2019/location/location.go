package location

type Location struct {
	X int
	Y int
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
	return abs(l.X-l2.X) + abs(l.Y-l2.Y)
}

func (l Location) Manhantan() int {
	return abs(l.X) + abs(l.Y)
}

func (l Location) Add(l2 Location) Location {
	return Location{l.X + l2.X, l.Y + l2.Y}
}

func (l Location) Direction(l2 Location) Location {
	x := l2.X - l.X
	y := l2.Y - l.Y
	d := gcd(x, y)

	return Location{x / d, y / d}
}

func (l Location) Mul(a int) Location {
	return Location{l.X * a, l.Y * a}
}

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
func (l Location) Distance(l2 Location) int {
	return abs(l.x-l2.x) + abs(l.y-l2.y)
}

func (l Location) Manhantan() int {
	return abs(l.x) + abs(l.y)
}

func (l Location) Add(l2 Location) Location {
	return Location{l.x+l2.x, l.y+l2.y}
}


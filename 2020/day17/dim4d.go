package main

type point4d struct {
	x, y, z, w int
}

var orig4d = point4d{}

var _ point = point4d{}

func (p point4d) Neighbors() []point {
	points := make([]point, 0, 80)
	for _, x := range []int{-1, 0, 1} {
		for _, y := range []int{-1, 0, 1} {
			for _, z := range []int{-1, 0, 1} {
				for _, w := range []int{-1, 0, 1} {
					b := point4d{x, y, z, w}
					if b != orig4d {
						points = append(points, p.Add(b))
					}
				}
			}
		}

	}
	return points
}

func (a point4d) Add(p point) point {
	b, ok := p.(point4d)
	if !ok {
		panic("Can not cast variable to point4d")
	}
	return point4d{a.x + b.x, a.y + b.y, a.z + b.z, a.w + b.w}
}

func (p point4d) AddScaler(i int) point {
	return point4d{p.x + i, p.y + i, p.z + i, p.w + i}
}

func (a point4d) Min(p point) point {
	b, ok := p.(point4d)
	if !ok {
		panic("Can not cast variable to point4d")
	}
	x := a.x
	y := a.y
	z := a.z
	w := a.w
	if b.x < x {
		x = b.x
	}
	if b.y < y {
		y = b.y
	}
	if b.z < z {
		z = b.z
	}
	if b.w < w {
		w = b.w
	}
	return point4d{x, y, z, w}
}
func (a point4d) Max(p point) point {
	b, ok := p.(point4d)
	if !ok {
		panic("Can not cast variable to point4d")
	}
	x := a.x
	y := a.y
	z := a.z
	w := a.w
	if b.x > x {
		x = b.x
	}
	if b.y > y {
		y = b.y
	}
	if b.z > z {
		z = b.z
	}
	if b.w > w {
		w = b.w
	}
	return point4d{x, y, z, w}
}

func (a point4d) Range(p point) []point {
	b, ok := p.(point4d)
	if !ok {
		panic("Can not cast variable to point4d")
	}
	out := []point{}
	for x := a.x; x < b.x+1; x++ {
		for y := a.y; y < b.y+1; y++ {
			for z := a.z; z < b.z+1; z++ {
				for w := a.w; w < b.w+1; w++ {
					out = append(out, point4d{x, y, z, w})
				}
			}
		}
	}
	return out
}

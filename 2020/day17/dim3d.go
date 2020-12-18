package main

type point3d struct {
	x, y, z int
}

var orig3d = point3d{}

var _ point = point3d{}

func (p point3d) Neighbors() []point {
	points := make([]point, 0, 26)
	for _, x := range []int{-1, 0, 1} {
		for _, y := range []int{-1, 0, 1} {
			for _, z := range []int{-1, 0, 1} {
				b := point3d{x, y, z}
				if b != orig3d {
					points = append(points, p.Add(b))
				}
			}
		}

	}
	return points
}

func (a point3d) Add(p point) point {
	b, ok := p.(point3d)
	if !ok {
		panic("Can not cast variable to point3d")
	}
	return point3d{a.x + b.x, a.y + b.y, a.z + b.z}
}

func (p point3d) AddScaler(i int) point {
	return point3d{p.x + i, p.y + i, p.z + i}
}

func (a point3d) Min(p point) point {
	b, ok := p.(point3d)
	if !ok {
		panic("Can not cast variable to point3d")
	}
	x := a.x
	y := a.y
	z := a.z
	if b.x < x {
		x = b.x
	}
	if b.y < y {
		y = b.y
	}
	if b.z < z {
		z = b.z
	}
	return point3d{x, y, z}
}
func (a point3d) Max(p point) point {
	b, ok := p.(point3d)
	if !ok {
		panic("Can not cast variable to point3d")
	}
	x := a.x
	y := a.y
	z := a.z
	if b.x > x {
		x = b.x
	}
	if b.y > y {
		y = b.y
	}
	if b.z > z {
		z = b.z
	}
	return point3d{x, y, z}
}

func (a point3d) Range(p point) []point {
	b, ok := p.(point3d)
	if !ok {
		panic("Can not cast variable to point3d")
	}
	out := []point{}
	for x := a.x; x < b.x+1; x++ {
		for y := a.y; y < b.y+1; y++ {
			for z := a.z; z < b.z+1; z++ {

				out = append(out, point3d{x, y, z})
			}
		}
	}
	return out
}

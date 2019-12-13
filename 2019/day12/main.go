package main

import (
	"fmt"
)

func main() {
	initial := []planet{
		//example1 179, 2772
		//{pos: point{-1,0,2}},
		//{pos: point{2,-10,-7}},
		//{pos: point{4,-8,8}},
		//{pos: point{3,5,-1}},
		//example2 1940, 4686774924
		//{pos: point{-8,-10,0}},
		//{pos: point{5,5,10}},
		//{pos: point{2,-7,3}},
		//{pos: point{9,-8,-3}},
		//input
		{pos: point{0, 4, 0}},
		{pos: point{-10, -6, -14}},
		{pos: point{9, -16, -3}},
		{pos: point{6, -1, 2}},
	}

	planets := make([]planet, 4)
	copy(planets, initial)

	for count := 0; count < 1000; count++ {

		for i := 0; i < len(planets)-1; i++ {
			for j := i + 1; j < len(planets); j++ {
				dp1, dp2 := gravity(planets[i], planets[j])

				planets[i].vel = add(dp1, planets[i].vel)
				planets[j].vel = add(dp2, planets[j].vel)
			}
		}
		energy := 0
		for i := range planets {
			planets[i].pos = add(planets[i].pos, planets[i].vel)
			energy += planets[i].vel.energy() * planets[i].pos.energy()
		}
		if count == 999 {
			fmt.Println(energy)
		}
	}

	count := 1

	copy(planets, initial)
	var repx, repy, repz int

	for repx == 0 || repy == 0 || repz == 0 {
		for i := 0; i < len(planets)-1; i++ {
			for j := i + 1; j < len(planets); j++ {
				dp1, dp2 := gravity(planets[i], planets[j])

				planets[i].vel = add(dp1, planets[i].vel)
				planets[j].vel = add(dp2, planets[j].vel)
			}
		}

		for i := range planets {
			planets[i].pos = add(planets[i].pos, planets[i].vel)
		}
		if repx == 0 && areSameX(initial, planets) {
			repx = count
		}
		if repy == 0 && areSameY(initial, planets) {
			repy = count
		}
		if repz == 0 && areSameZ(initial, planets) {
			repz = count
		}

		count++

	}

	fmt.Println(lcm(repx, lcm(repy, repz)))

}

func areSameX(initial, planets []planet) bool {
	return initial[0].pos.x == planets[0].pos.x && initial[0].vel.x == planets[0].vel.x &&
		initial[1].pos.x == planets[1].pos.x && initial[1].vel.x == planets[1].vel.x &&
		initial[2].pos.x == planets[2].pos.x && initial[2].vel.x == planets[2].vel.x &&
		initial[3].pos.x == planets[3].pos.x && initial[3].vel.x == planets[3].vel.x
}
func areSameY(initial, planets []planet) bool {
	return initial[0].pos.y == planets[0].pos.y && initial[0].vel.y == planets[0].vel.y &&
		initial[1].pos.y == planets[1].pos.y && initial[1].vel.y == planets[1].vel.y &&
		initial[2].pos.y == planets[2].pos.y && initial[2].vel.y == planets[2].vel.y &&
		initial[3].pos.y == planets[3].pos.y && initial[3].vel.y == planets[3].vel.y
}
func areSameZ(initial, planets []planet) bool {
	return initial[0].pos.z == planets[0].pos.z && initial[0].vel.z == planets[0].vel.z &&
		initial[1].pos.z == planets[1].pos.z && initial[1].vel.z == planets[1].vel.z &&
		initial[2].pos.z == planets[2].pos.z && initial[2].vel.z == planets[2].vel.z &&
		initial[3].pos.z == planets[3].pos.z && initial[3].vel.z == planets[3].vel.z
}

type point struct {
	x, y, z int
}

func (p point) energy() int {
	return abs(p.x) + abs(p.y) + abs(p.z)
}

func add(p1, p2 point) point {
	p := point{}
	p.x = p1.x + p2.x
	p.y = p1.y + p2.y
	p.z = p1.z + p2.z
	return p
}

type planet struct {
	pos point
	vel point
}

func gravity(p1, p2 planet) (point, point) {
	var v1, v2 point
	v1.x = diff(p1.pos.x, p2.pos.x)
	v2.x = v1.x * -1
	v1.y = diff(p1.pos.y, p2.pos.y)
	v2.y = v1.y * -1
	v1.z = diff(p1.pos.z, p2.pos.z)
	v2.z = v1.z * -1
	return v1, v2
}

func diff(x, y int) int {
	if x > y {
		return -1
	}
	if x < y {
		return 1
	}
	return 0

}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func gcd(a, b int) int {
	a = abs(a)
	b = abs(b)

	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

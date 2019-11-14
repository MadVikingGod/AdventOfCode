package main

import (
	"fmt"
	"github.com/madvikinggod/AdventOfCode/2018/helpers"
	"log"
)

func main() {
	input, err := helpers.GetInput(15)
	if err != nil {
		log.Panic(err)
	}
	gameMap := readMap(input)
	fmt.Print(gameMap)
	for i := 0; gameStep(gameMap); i++ {
		fmt.Print(gameMap)
		fmt.Printf("%#v\n", gameMap.goblins)
		if i > 50 {
			fmt.Println("OH NOES!")
			break
		}
	}
	fmt.Printf("%#v\n", gameMap.goblins)

}

func gameStep(m *gameMap) bool {
	moved := map[location]bool{}
	for x := 0; x < m.x; x++ {
		for y := 0; y < m.y; y++ {

			loc := location{x, y}
			if moved[loc] {
				continue
			}

			var enemies map[location]mapSquare
			var attackPower int
			if e, ok := m.elfs[loc]; ok {
				enemies = m.goblins
				attackPower = e.attackPower
			} else if g, ok := m.goblins[loc]; ok {
				enemies = m.elfs
				attackPower = g.attackPower
			} else {
				continue
			}

			if len(enemies) == 0 {
				return false
			}
			dst, move := findStep(*m, loc, enemies)
			if move {
				m.move(loc, dst)
			}
			moved[dst] = true

			//Attack

			for _, a := range dst.Adjecent() {
				if _, ok := enemies[a]; ok {
					m.attack(a, attackPower)
					continue
				}
			}
		}
	}
	return true
}

func findAdj(m gameMap, locs map[location]mapSquare) map[location]bool {
	adjSpaces := map[location]bool{}
	for l := range locs {
		adj := l.Adjecent()
		for _, a := range adj {
			if !m.isOccupied(a) {
				adjSpaces[a] = true
			}
		}
	}
	return adjSpaces
}

func findStep(m gameMap, start location, enemies map[location]mapSquare) (location, bool) {
	// Don't move if already adjecent
	for _, a := range start.Adjecent() {
		if _, ok := enemies[a]; ok {
			return start, false
		}
	}

	destinations := findAdj(m, enemies)
	type direction struct {
		loc   location
		start location
	}
	visited := map[location]bool{}
	queue := []direction{}

	for _, a := range start.Adjecent() {
		if !m.isOccupied(a) {
			queue = append(queue, direction{a, a})
		}
	}

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]
		visited[current.loc] = true
		if destinations[current.loc] {
			return current.start, true
		}
		if start.x == 4 && start.y == 4 {
			fmt.Println(current.loc, queue)
		}
		for _, a := range current.loc.Adjecent() {

			if !m.isOccupied(a) && !visited[a] {
				queue = append(queue, direction{a, current.start})
			}
		}
	}

	return start, false
}

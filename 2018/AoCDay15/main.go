package main

import (
	"fmt"
	"log"

	"github.com/madvikinggod/AdventOfCode/2018/helpers"
)

func main() {
	input, err := helpers.GetInput(15)
	if err != nil {
		log.Panic(err)
	}
	gameMap := readMap(input)
	fmt.Print(gameMap)
	n := runGame(gameMap)
	println(n, "*", gameMap.getHealth(), "=", n*gameMap.getHealth())

	attackMap := readMap(input)
	atkPower, rounds, health := findElfPower(attackMap)
	fmt.Println(atkPower)
	println(rounds, "*", health, "=", n*health)

}

func findElfPower(m *gameMap) (int, int, int) {
	origMap := *m
	attackPower := 4
	rounds := 0
	health := 0
	for ; attackPower < 30; attackPower++ {
		currentMap := origMap
		currentMap.elfs = map[location]mapSquare{}
		currentMap.goblins = map[location]mapSquare{}
		for loc, g := range origMap.goblins {
			currentMap.goblins[loc] = g
		}
		for loc, e := range origMap.elfs {
			e.attackPower = attackPower
			currentMap.elfs[loc] = e
		}
		rounds = runGame(&currentMap)
		health = currentMap.getHealth()
		fmt.Printf("%d, %d: %d * %d = %d\n", attackPower, len(currentMap.elfs), rounds, health, rounds*health)

	}

	return attackPower, rounds, health
}

func runGame(m *gameMap) int {
	i := 0
	for ; gameStep(m); i++ {
		if i > 500 {
			fmt.Println("OH NOES!")
			break
		}

	}
	return i
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

			enemy, attack := findTarget(dst, enemies)
			if attack {
				m.attack(enemy, attackPower)
			}
		}
	}
	return true
}

func findTarget(loc location, enemies map[location]mapSquare) (location, bool) {
	targetHP := 0
	var targetLoc location
	found := false
	for _, a := range loc.Adjecent() {
		if e, ok := enemies[a]; ok && (found == false || e.hitpoints < targetHP) {
			targetHP = e.hitpoints
			targetLoc = a
			found = true
		}
	}
	return targetLoc, found
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
		for _, a := range current.loc.Adjecent() {

			if !m.isOccupied(a) && !visited[a] {
				queue = append(queue, direction{a, current.start})
				visited[a] = true
			}
		}
	}

	return start, false
}

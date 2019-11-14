package main

import (
	"fmt"
	"strings"
)

func readMap(input []string) *gameMap {

	out := NewGameMap(len(input), len(input[0]))
	for i, line := range input {
		for j, c := range line {
			loc := location{i, j}
			square := NewMapSquare(spaces[c])
			switch square.kind {
			case wall:
				out.walls[loc] = square
			case elf:
				out.elfs[loc] = square
			case goblin:
				out.goblins[loc] = square
			}
		}
		i++
	}
	return out
}

type location struct {
	x int
	y int
}

func (l location) Adjecent() []location {
	return []location{
		location{l.x - 1, l.y}, // up
		location{l.x, l.y - 1}, // left
		location{l.x, l.y + 1}, // right
		location{l.x + 1, l.y}, // down
	}
}

const (
	wall = iota
	empty
	elf
	goblin
)

var spaces = map[rune]int{
	'#': wall,
	'.': empty,
	'E': elf,
	'G': goblin,
}
var invSpaces = []string{"#", ".", "E", "G"}

type gameMap struct {
	x       int
	y       int
	walls   map[location]mapSquare
	goblins map[location]mapSquare
	elfs    map[location]mapSquare
}

func NewGameMap(x, y int) *gameMap {
	return &gameMap{
		x:       x,
		y:       y,
		walls:   map[location]mapSquare{},
		goblins: map[location]mapSquare{},
		elfs:    map[location]mapSquare{},
	}
}

func (m gameMap) GetString(loc location) string {
	if s, ok := m.walls[loc]; ok {
		return invSpaces[s.kind]
	}
	if s, ok := m.elfs[loc]; ok {
		return invSpaces[s.kind]
	}
	if s, ok := m.goblins[loc]; ok {
		return invSpaces[s.kind]
	}
	return "."
}

func (m gameMap) getUnits(loc []location) []string {
	units := []string{}
	for _, l := range loc {
		if e, ok := m.elfs[l]; ok {
			units = append(units, e.GoString())
		}
		if g, ok := m.goblins[l]; ok {
			units = append(units, g.GoString())
		}
	}
	return units
}

func (m gameMap) String() string {
	b := strings.Builder{}
	for x := 0; x < m.x; x++ {
		units := []location{}
		for y := 0; y < m.y; y++ {
			s := m.GetString(location{x, y})
			if s == "E" || s == "G" {
				units = append(units, location{x, y})
			}
			b.WriteString(s)
		}
		if len(units) > 0 {
			b.WriteString("\t{")
			b.WriteString(strings.Join(m.getUnits(units), "}, {"))
			b.WriteString("}")
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (m gameMap) isOccupied(loc location) bool {
	validX := loc.x >= 0 && loc.x < m.x
	validY := loc.y >= 0 && loc.y < m.y
	if !validX || !validY {
		return true
	}
	_, wok := m.walls[loc]
	_, eok := m.elfs[loc]
	_, gok := m.goblins[loc]
	return wok || eok || gok
}

func (m *gameMap) move(start, dest location) {
	if e, ok := m.elfs[start]; ok {
		m.elfs[dest] = e
		delete(m.elfs, start)
	}
	if g, ok := m.goblins[start]; ok {
		m.goblins[dest] = g
		delete(m.goblins, start)
	}
}

func (m *gameMap) attack(loc location, dmg int) {

	if e, ok := m.elfs[loc]; ok {
		e.hitpoints -= dmg
		m.elfs[loc] = e
		if e.hitpoints <= 0 {
			delete(m.elfs, loc)
		}
	}
	if g, ok := m.goblins[loc]; ok {
		g.hitpoints -= dmg
		m.goblins[loc] = g
		if g.hitpoints <= 0 {
			delete(m.goblins, loc)
		}
	}
}

func (m gameMap) getHealth() int {
	health := 0
	for _, e := range m.elfs {
		health += e.hitpoints
	}
	for _, g := range m.goblins {
		health += g.hitpoints
	}
	return health
}

type mapSquare struct {
	kind        int
	hitpoints   int
	attackPower int
}

func NewMapSquare(kind int) mapSquare {
	return mapSquare{
		kind:        kind,
		hitpoints:   200,
		attackPower: 3,
	}
}
func (m mapSquare) String() string {
	return invSpaces[m.kind]
}

func (m mapSquare) GoString() string {
	return fmt.Sprintf("%s %3d %d", invSpaces[m.kind], m.hitpoints, m.attackPower)
}

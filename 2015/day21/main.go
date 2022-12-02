package main

func main() {
	minCost := 1000000

	for _, weapon := range weapons {
		for _, armor := range armors {
			for _, ring1 := range rings {
				for _, ring2 := range rings {
					if ring1 == ring2 {
						continue
					}

					cost := weapon.cost + armor.cost + ring1.cost + ring2.cost
					if cost >= minCost {
						continue
					}
					var boss = &sprite{104, 8, 1}

					if canWin(boss, newPlayer(weapon, armor, ring1, ring2)) {
						minCost = min(minCost, cost)
					}
				}
			}
		}
	}
	println(minCost)

	maxCost := 0
	for _, weapon := range weapons {
		for _, armor := range armors {
			for _, ring1 := range rings {
				for _, ring2 := range rings {
					if ring1 == ring2 {
						continue
					}

					cost := weapon.cost + armor.cost + ring1.cost + ring2.cost
					if cost < maxCost {
						continue
					}
					var boss = &sprite{104, 8, 1}

					if !canWin(boss, newPlayer(weapon, armor, ring1, ring2)) {
						maxCost = max(maxCost, cost)
					}
				}
			}
		}
	}
	println(maxCost)
}

func canWin(boss, s *sprite) bool {
	for {
		won, lost := s.turn(boss)
		if won {
			return true
		}
		if lost {
			return false
		}
	}
}

type sprite struct {
	hp, damage, armor int
}

func newPlayer(weapon, armor, ring1, ring2 item) *sprite {
	return &sprite{
		hp:     100,
		damage: weapon.damage + ring1.damage + ring2.damage,
		armor:  armor.armor + ring1.armor + ring2.armor,
	}
}

func (s *sprite) turn(e *sprite) (bool, bool) {
	e.hp -= max(1, s.damage-e.armor)
	if e.hp <= 0 {
		return true, false
	}
	s.hp -= max(1, e.damage-s.armor)
	return false, s.hp <= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type item struct {
	name   string
	cost   int
	damage int
	armor  int
}

/*
Weapons:    Cost  Damage  Armor
Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0

Armor:      Cost  Damage  Armor
Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5

Rings:      Cost  Damage  Armor
Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3
*/

var weapons = []item{
	{"Dagger", 8, 4, 0},
	{"Shortsword", 10, 5, 0},
	{"Warhammer", 25, 6, 0},
	{"Longsword", 40, 7, 0},
	{"Greataxe", 74, 8, 0},
}

var armors = []item{
	{"none", 0, 0, 0},
	{"Leather", 13, 0, 1},
	{"Chainmail", 31, 0, 2},
	{"Splintmail", 53, 0, 3},
	{"Bandedmail", 75, 0, 4},
	{"Platemail", 102, 0, 5},
}

var rings = []item{
	{"none", 0, 0, 0},
	{"none2", 0, 0, 0},
	{"Damage +1", 25, 1, 0},
	{"Damage +2", 50, 2, 0},
	{"Damage +3", 100, 3, 0},
	{"Defense +1", 20, 0, 1},
	{"Defense +2", 40, 0, 2},
	{"Defense +3", 80, 0, 3},
}

package main

import _ "embed"

//go:embed input.txt
var input string

func main() {
	queue := []state{{pHP: 50, pMana: 500, bHP: 58, bDmg: 9}}
	// queue := []state{{pHP: 10, pMana: 250, bHP: 13, bDmg: 8}}
	minMana := 1<<31 - 1
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		if minMana < s.mana {
			continue
		}
		if s.pMana < 53 && s.recharge == 0 {
			continue
		}
		for i, sp := range spells {
			switch i {
			case 2:
				if s.poison > 0 {
					continue
				}
			case 3:
				if s.shield > 0 {
					continue
				}
			case 4:
				if s.recharge > 0 {
					continue
				}
			}
			if s.pMana < costs[i] {
				continue
			}
			s2, win, lose := s.turn(sp)
			if win && s2.mana < minMana {
				minMana = s2.mana
				continue
			}
			if !lose {
				queue = append(queue, s2)
			}

		}

	}
	println(minMana)
}

type state struct {
	pHP, pMana int
	bHP, bDmg  int

	// turns left
	shield, poison, recharge int

	// mana spent
	mana int
}

type spell func(*state)

var costs = []int{53, 73, 173, 113, 229}

var spells = []spell{
	func(s *state) {
		s.pMana -= 53
		s.bHP -= 4
		s.mana += 53
	},
	func(s *state) {
		s.pMana -= 73
		s.bHP -= 2
		s.pHP += 2
		s.mana += 73
	},
	func(s *state) {
		s.pMana -= 113
		s.shield = 6
		s.mana += 113
	},
	func(s *state) {
		s.pMana -= 173
		s.poison = 6
		s.mana += 173
	},
	func(s *state) {
		s.pMana -= 229
		s.recharge = 5
		s.mana += 229
	},
}

func (s state) turn(sp spell) (state, bool, bool) {
	// player turn
	if s.poison > 0 {
		s.bHP -= 3
		s.poison--
	}
	if s.bHP <= 0 {
		return s, true, false
	}
	if s.recharge > 0 {
		s.pMana += 101
		s.recharge--
	}
	if s.shield > 0 {
		s.shield--
	}
	sp(&s)

	// boss turn
	if s.poison > 0 {
		s.bHP -= 3
		s.poison--
	}
	if s.bHP <= 0 {
		return s, true, false
	}
	if s.recharge > 0 {
		s.pMana += 101
		s.recharge--
	}
	var shield int
	if s.shield > 0 {
		shield = 7
		s.shield--
	}
	s.pHP -= max(1, s.bDmg-shield)
	return s, false, s.pHP <= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

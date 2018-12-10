package main

import (
	"fmt"
)

var (
	player     = make([]int, 478)
	lastMarble = 71240
)

func main() {
	//`478 players; last marble is worth 71240 points`
	player = make([]int, 478)
	lastMarble := 7124000
	board := newBoard()
	fmt.Println(board.placed)
	currentPlayer := 3
	var add int
	for i := 5; i < lastMarble; i++ {
		currentPlayer++
		if i%23 == 0 {
			score := player[currentPlayer%len(player)]
			add = i + board.sevenBack()
			score = score + add
			player[currentPlayer%len(player)] = score
			// fmt.Println(i, board.placed)

			continue
		}

		board.Add(i)
	}

	fmt.Println(add)
	fmt.Println(player)
	fmt.Println(max(player))

}

func max(a []int) int {
	m := 0
	for _, x := range a {
		if x > m {
			m = x
		}
	}
	return m
}

type Board struct {
	lastMarble int
	placed     []int
}

func newBoard() *Board {
	return &Board{
		lastMarble: 1,
		placed:     []int{0, 4, 2, 1, 3},
	}
}

func (b *Board) Add(i int) {
	b.placed = append(b.placed, 0)
	newLoc := (b.lastMarble + 2) % (len(b.placed))
	if newLoc == 0 {
		newLoc++
	}
	copy(b.placed[newLoc+1:], b.placed[newLoc:])
	b.placed[newLoc] = i
	b.lastMarble = newLoc
}

func (b *Board) sevenBack() int {
	loc := b.lastMarble - 7

	if loc < 0 {
		loc = loc + len(b.placed)
	}
	s := b.placed[loc]
	b.placed = append(b.placed[:loc], b.placed[loc+1:]...)
	b.lastMarble = loc
	return s
}

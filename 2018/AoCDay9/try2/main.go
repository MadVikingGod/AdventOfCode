package main

import "fmt"

//`478 players; last marble is worth 71240 points`
//10 players; last marble is worth 1618 points: high score is 8317
//13 players; last marble is worth 7999 points: high score is 146373
// 17 players; last marble is worth 1104 points: high score is 2764
// 21 players; last marble is worth 6111 points: high score is 54718
// 30 players; last marble is worth 5807 points: high score is 37305
var (
	players    int = 10
	lastMarble int = 1618
)

func main() {
	fmt.Println(playGame(9, 25))
	fmt.Println(playGame(10, 1618) == 8317)
	fmt.Println(playGame(13, 7999) == 146373)
	fmt.Println(playGame(478, 7124000))
}

func playGame(players, lastMarble int) (maxScore int) {
	current := &marble{0, nil, nil}
	current.prev = current
	current.next = current

	marble := 1
	scores := make([]int, players)
	for marble <= lastMarble {
		for elf := 0; elf < players && marble <= lastMarble; elf++ {
			if marble%23 == 0 {

				scores[elf] += marble
				removed := current.prev.prev.prev.prev.prev.prev.prev
				scores[elf] += removed.number
				removed.remove()

				current = removed.next
			} else {
				current = current.next
				current.insert(marble)
				current = current.next
			}
			marble++
			//root.print()
		}
	}

	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}

type marble struct {
	number int
	prev   *marble
	next   *marble
}

func (m *marble) print() {
	fmt.Printf("[%d ", m.number)
	current := m
	for current.next != m {
		current = current.next
		fmt.Printf("%d ", current.number)
	}
	fmt.Println("]")
}

func (m *marble) verify() {
	prev := m
	current := m.next
	for current != m {
		if prev.next != current {
			fmt.Printf("forward chain broken %d\n", prev.number)
		}
		if current.prev != prev {
			fmt.Printf("backward chain broken %d\n", current.number)
		}
		prev = current
		current = current.next
	}
}

func (m *marble) insert(i int) {
	newMarble := &marble{
		number: i,
		prev:   m,
		next:   m.next,
	}
	m.next.prev = newMarble
	m.next = newMarble

}
func (m *marble) remove() {
	m.next.prev = m.prev
	m.prev.next = m.next
}

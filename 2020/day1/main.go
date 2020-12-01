package main

import "fmt"

func main() {
	fmt.Println("Part1: ", SolvePart1(inputs))
	fmt.Println("Part2: ", SolvePart2(inputs))
}
func SolvePart1(inputs []int) int {
	s := set{}
	for _, i := range inputs {
		j := 2020 - i
		if s.Has(j) {
			return j * i
		}
		s.Add(i)
	}

	return 0
}

func SolvePart2(inputs []int) int {
	s := set{}
	s.Add(inputs[0])
	for i := 0; i < len(inputs)-1; i++ {
		x := inputs[i]
		for j := 0; j < len(inputs); j++ {
			y := inputs[j]
			s.Add(y)
			z := 2020 - x - y
			if s.Has(z) {
				return x * y * z
			}
		}
	}
	return 0
}

type set map[int]struct{}

func (s set) Add(i int) {
	s[i] = struct{}{}
}

func (s set) Delete(i int) {
	delete(s, i)
}

func (s set) Has(i int) bool {
	_, ok := s[i]
	return ok
}

package main

import "fmt"

func main() {
	set := new(25)
	out := 0
	for _, i := range input {
		if !set.isValid(i) && i != 127 {
			fmt.Println(i)
			out = i
			break
		}
		set.add(i)
	}

	list := find(out, input)
	min, max := minMax(list)
	fmt.Println(min + max)
}

func minMax(l []int) (int, int) {
	min := l[0]
	max := l[0]
	for _, i := range l {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	return min, max
}

type LimitedSet struct {
	size  int
	items map[int]int
	stack []int
}

func find(target int, l []int) []int {
	for i := range l {
		if out := lookahead(target, i, l); out != nil {
			return out
		}
	}
	return nil
}

func lookahead(target, start int, l []int) []int {
	if start > len(l) {
		return nil
	}
	sum := 0
	for i := start; i < len(l); i++ {
		sum += l[i]
		if sum == target {
			return l[start : i+1]
		}
		if sum > target {
			return nil
		}
	}
	return nil
}

func new(size int) LimitedSet {
	return LimitedSet{
		size:  size,
		items: map[int]int{},
		stack: make([]int, 0, size+1),
	}
}

func (s *LimitedSet) add(i int) {
	s.items[i] += 1
	s.stack = append(s.stack, i)
	if len(s.stack) > s.size {
		s.del(s.stack[0])
		s.stack = s.stack[1:]
	}
}
func (s *LimitedSet) del(i int) {
	s.items[i] -= 1
}
func (s *LimitedSet) isValid(i int) bool {
	if len(s.stack) < s.size {
		return true
	}
	for _, h := range s.stack {
		if s.items[i-h] > 0 {
			return true
		}
	}
	return false
}

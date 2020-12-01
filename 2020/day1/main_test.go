package main

import "testing"

func Test_SolvePar1(t *testing.T) {
	inputs := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	got := SolvePart1(inputs)
	if got != 514579 {
		t.Errorf("Part1 returned the wrong anser, got = %d, wanted = %d", got, 514579)
	}
}
func Test_SolvePar2(t *testing.T) {
	inputs := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	got := SolvePart2(inputs)
	if got != 241861950 {
		t.Errorf("Part1 returned the wrong anser, got = %d, wanted = %d", got, 241861950)
	}
}

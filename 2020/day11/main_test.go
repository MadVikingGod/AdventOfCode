package main

import (
	"fmt"
	"testing"
)

func TestFerry_Next(t *testing.T) {
	var (
		testSrc = []string{
			"L.LL.LL.LL",
			"LLLLLLL.LL",
			"L.L.L..L..",
			"LLLL.LL.LL",
			"L.LL.LL.LL",
			"L.LLLLL.LL",
			"..L.L.....",
			"LLLLLLLLLL",
			"L.LLLLLL.L",
			"L.LLLLL.LL",
		}
		testStep1 = []string{
			"#.##.##.##",
			"#######.##",
			"#.#.#..#..",
			"####.##.##",
			"#.##.##.##",
			"#.#####.##",
			"..#.#.....",
			"##########",
			"#.######.#",
			"#.#####.##",
		}
		testStep2 = []string{
			"#.LL.L#.##",
			"#LLLLLL.L#",
			"L.L.L..L..",
			"#LLL.LL.L#",
			"#.LL.LL.LL",
			"#.LLLL#.##",
			"..L.L.....",
			"#LLLLLLLL#",
			"#.LLLLLL.L",
			"#.#LLLL.##",
		}
		testStep3 = []string{
			"#.##.L#.##",
			"#L###LL.L#",
			"L.#.#..#..",
			"#L##.##.L#",
			"#.##.LL.LL",
			"#.###L#.##",
			"..#.#.....",
			"#L######L#",
			"#.LL###L.L",
			"#.#L###.##",
		}
		testStep4 = []string{
			"#.#L.L#.##",
			"#LLL#LL.L#",
			"L.L.L..#..",
			"#LLL.##.L#",
			"#.LL.LL.LL",
			"#.LL#L#.##",
			"..L.L.....",
			"#L#LLLL#L#",
			"#.LLLLLL.L",
			"#.#L#L#.##",
		}
		testStep5 = []string{
			"#.#L.L#.##",
			"#LLL#LL.L#",
			"L.#.L..#..",
			"#L##.##.L#",
			"#.#L.LL.LL",
			"#.#L#L#.##",
			"..L.L.....",
			"#L#L##L#L#",
			"#.LLLLLL.L",
			"#.#L#L#.##",
		}
	)
	current := new(testSrc)
	nxt := &Ferry{seats: map[point]seat{}}

	current, nxt = current.Next(nxt)
	if !current.Equal(new(testStep1)) {
		t.Error("Step 1 didn't match")
		t.FailNow()
	}
	current, nxt = current.Next(nxt)
	if !current.Equal(new(testStep2)) {
		t.Error("Step 2 didn't match")
		t.FailNow()
	}
	current, nxt = current.Next(nxt)
	if !current.Equal(new(testStep3)) {
		t.Error("Step 3 didn't match")
		t.FailNow()
	}
	current, nxt = current.Next(nxt)
	if !current.Equal(new(testStep4)) {
		t.Error("Step 4 didn't match")
		t.FailNow()
	}
	current, nxt = current.Next(nxt)
	if !current.Equal(new(testStep5)) {
		t.Error("Step 5 didn't match")
	}

}

func TestFerry_Next2(t *testing.T) {
	var (
		testSrc = []string{
			"L.LL.LL.LL",
			"LLLLLLL.LL",
			"L.L.L..L..",
			"LLLL.LL.LL",
			"L.LL.LL.LL",
			"L.LLLLL.LL",
			"..L.L.....",
			"LLLLLLLLLL",
			"L.LLLLLL.L",
			"L.LLLLL.LL",
		}
		testStep1 = []string{
			"#.##.##.##",
			"#######.##",
			"#.#.#..#..",
			"####.##.##",
			"#.##.##.##",
			"#.#####.##",
			"..#.#.....",
			"##########",
			"#.######.#",
			"#.#####.##",
		}
		testStep2 = []string{
			"#.LL.LL.L#",
			"#LLLLLL.LL",
			"L.L.L..L..",
			"LLLL.LL.LL",
			"L.LL.LL.LL",
			"L.LLLLL.LL",
			"..L.L.....",
			"LLLLLLLLL#",
			"#.LLLLLL.L",
			"#.LLLLL.L#",
		}
		testStep3 = []string{
			"#.L#.##.L#",
			"#L#####.LL",
			"L.#.#..#..",
			"##L#.##.##",
			"#.##.#L.##",
			"#.#####.#L",
			"..#.#.....",
			"LLL####LL#",
			"#.L#####.L",
			"#.L####.L#",
		}
		testStep4 = []string{
			"#.L#.L#.L#",
			"#LLLLLL.LL",
			"L.L.L..#..",
			"##LL.LL.L#",
			"L.LL.LL.L#",
			"#.LLLLL.LL",
			"..L.L.....",
			"LLLLLLLLL#",
			"#.LLLLL#.L",
			"#.L#LL#.L#",
		}
		testStep5 = []string{
			"#.L#.L#.L#",
			"#LLLLLL.LL",
			"L.L.L..#..",
			"##L#.#L.L#",
			"L.L#.#L.L#",
			"#.L####.LL",
			"..#.#.....",
			"LLL###LLL#",
			"#.LLLLL#.L",
			"#.L#LL#.L#",
		}
		testStep6 = []string{
			"#.L#.L#.L#",
			"#LLLLLL.LL",
			"L.L.L..#..",
			"##L#.#L.L#",
			"L.L#.LL.L#",
			"#.LLLL#.LL",
			"..#.L.....",
			"LLL###LLL#",
			"#.LLLLL#.L",
			"#.L#LL#.L#",
		}
	)
	current := new(testSrc)
	nxt := &Ferry{seats: map[point]seat{}}
	current, nxt = current.Next2(nxt)
	if !current.Equal(new(testStep1)) {
		t.Error("Step 1 didn't match")
		t.FailNow()
	}
	fmt.Println(current)
	fmt.Println("--------------")
	current, nxt = current.Next2(nxt)
	if !current.Equal(new(testStep2)) {
		t.Error("Step 2 didn't match")
		fmt.Println(current)
		t.FailNow()
	}
	current, nxt = current.Next2(nxt)
	if !current.Equal(new(testStep3)) {
		t.Error("Step 3 didn't match")
		t.FailNow()
	}
	current, nxt = current.Next2(nxt)
	if !current.Equal(new(testStep4)) {
		t.Error("Step 4 didn't match")
		t.FailNow()
	}
	current, nxt = current.Next2(nxt)
	if !current.Equal(new(testStep5)) {
		t.Error("Step 5 didn't match")
	}
	current, nxt = current.Next2(nxt)
	if !current.Equal(new(testStep6)) {
		t.Error("Step 6 didn't match")
	}
}
